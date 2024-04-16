package model

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"strings"
)

type clickHouseEntityPropertyRepository struct {
	conn driver.Conn
}

func NewClickHouseEntityPropertyRepository(conn driver.Conn) EntityPropertyRepository {
	return &clickHouseEntityPropertyRepository{conn: conn}
}

func (repo *clickHouseEntityPropertyRepository) Save(properties []EntityProperty) {

	batch := repo.CreateBatch()
	repo.Add(batch, properties)
	repo.SendBatch(batch)
}

func (repo *clickHouseEntityPropertyRepository) CreateBatch() driver.Batch {
	batch, err := repo.conn.PrepareBatch(context.Background(), "INSERT INTO entity")
	if err != nil {
		panic(err)
	}

	return batch
}

func (repo *clickHouseEntityPropertyRepository) SendBatch(batch driver.Batch) {
	err := batch.Send()
	if err != nil {
		panic(err)
	}
}

func (repo *clickHouseEntityPropertyRepository) Add(batch driver.Batch, properties []EntityProperty) {

	for _, prop := range properties {
		propStruct := prop.(*entityProperty)
		if propStruct.state.isDirty {
			//log.Println(fmt.Sprintf("Saving %v=%v", prop.Name(), prop.Value()))
			err := batch.AppendStruct(propStruct.state)
			if err != nil {
				panic(err)
			}
		}
	}
}

func getPaginatorSql(offset int64, totalRows int64, sortField string, sortFieldType string, sortOrder bool) string {
	filtrOrderSql := getFiltrOrderSql(false, sortField, sortFieldType, sortOrder)

	return fmt.Sprintf("%v LIMIT %v OFFSET %v", filtrOrderSql, totalRows, offset)
}

var fieldTypeMap = map[string]string{
	"int":      "PropertyIntValue",
	"string":   "PropertyStringValue",
	"float":    "PropertyFloatValue",
	"bool":     "PropertyBoolValue",
	"datetime": "PropertyDateTimeValue",
}

func getFiltrOrderSql(cnt bool, sortField string, fieldType string, sortOrder bool) string {

	selectPart := "filterSql.EntityId"
	strOrderSqlPart := ""
	if cnt {
		selectPart = "count(1)"
	} else {
		if sortOrder {
			if sortField == "EntityId" {
				strOrderSqlPart = "ORDER BY filterSql.EntityId DESC"
			} else {
				strOrderSqlPart = "ORDER BY filterSql.AggPropertyValue DESC"
			}
		} else {
			if sortField == "EntityId" {
				strOrderSqlPart = "ORDER BY filterSql.EntityId ASC"
			} else {
				strOrderSqlPart = "ORDER BY filterSql.AggPropertyValue ASC"
			}
		}
	}

	if sortField == "EntityId" {
		return fmt.Sprintf(`SELECT %v FROM (SELECT EntityId FROM entity GROUP BY (EntityId)) as filterSql %v`,
			selectPart, strOrderSqlPart)

	}

	valueField := fieldTypeMap[fieldType]

	sql := fmt.Sprintf(`SELECT %v
FROM (
         SELECT EntityId,
                argMax(entity.%v, entity.Date) as AggPropertyValue,
                argMax(IsNull, entity.Date)  as AggIsNull
         FROM entity
         WHERE PropertyName == '%v'
         GROUP BY (EntityId, PropertyName)
         HAVING AggIsNull == false
) as filterSql %v`, selectPart, valueField, sortField, strOrderSqlPart)

	return sql
}

func (repo *clickHouseEntityPropertyRepository) Find(whereAndPart []string, offset int64, totalRows int64, sortField string, sortFieldType string, sortOrder bool) EntityPropertyIterator {

	fields := []string{
		"PropertyType",
		"PropertyStringValue",
		"PropertyIntValue",
		"PropertyFloatValue",
		"PropertyBoolValue",
		"PropertyDateTimeValue",
		"IsNull",
	}
	sqlFields := make([]string, len(fields))

	for i := range fields {
		sqlFields[i] = "argMax(" + fields[i] + ", entity.Date) as " + fields[i]
	}

	sqlStaticFields := []string{"EntityId", "PropertyName"}

	selectPart := "SELECT " + strings.Join(append(sqlFields, sqlStaticFields...), ",")
	fromPart := "FROM entity"

	paginatorSql := getPaginatorSql(offset, totalRows, sortField, sortFieldType, sortOrder)

	usePaginatorSql := false
	if len(whereAndPart) == 0 {
		usePaginatorSql = true
		whereAndPart = append(whereAndPart, "EntityId IN ("+paginatorSql+")")
	}

	var wherePart string
	if len(whereAndPart) > 0 {
		wherePart = "WHERE " + strings.Join(whereAndPart, " AND ")
	}

	groupPart := "GROUP BY (EntityId, PropertyName) HAVING IsNull == false"

	sql := strings.Join([]string{selectPart, fromPart, wherePart, groupPart}, " ")
	if usePaginatorSql {
		finalFields := sqlStaticFields
		for _, field := range fields {
			finalFields = append(finalFields, field)
		}

		sql = strings.Join([]string{
			"SELECT " + strings.Join(finalFields, ",") + " FROM",
			"(" + paginatorSql + ") as orderedIds",
			"JOIN",
			"(" + sql + ") as filteredProps",
			"ON orderedIds.EntityId = filteredProps.EntityId"}, " ")
	}

	log.Println(sql)

	rows, err := repo.conn.Query(context.Background(), sql)
	if err != nil {
		panic(err)
	}

	totalSql := getFiltrOrderSql(true, sortField, sortFieldType, sortOrder)
	log.Println(totalSql)

	row := repo.conn.QueryRow(context.Background(), totalSql)
	var total uint64
	err = row.Scan(&total)
	if err != nil {
		panic(err)
	}

	return &rowsIterator{rows: rows, total: total}
}

func (repo *clickHouseEntityPropertyRepository) FindAllUniqueNames() []string {

	sql := `SELECT PropertyName FROM entity GROUP BY PropertyName`

	log.Println(sql)

	rows, err := repo.conn.Query(context.Background(), sql)
	if err != nil {
		panic(err)
	}

	var propName string
	var propNames []string = make([]string, 0, 20)

	for rows.Next() {
		err = rows.Scan(&propName)
		if err != nil {
			panic(err)
		}
		if propName != "" {
			propNames = append(propNames, propName)
		}
	}

	return propNames
}

type rowsIterator struct {
	rows  driver.Rows
	total uint64
}

func (rI *rowsIterator) Total() uint64 {
	return rI.total
}

func (rI *rowsIterator) Next() bool {
	isNext := rI.rows.Next()
	if !isNext {
		err := rI.rows.Close()
		if err != nil {
			panic(err)
		}
	}

	return isNext
}

func (rI *rowsIterator) Get() EntityProperty {
	var rowStruct entityProperty = entityProperty{state: &entityPropertyState{}}

	err := rI.rows.ScanStruct(rowStruct.state)
	if err != nil {
		panic(err)
	}

	return &rowStruct
}
