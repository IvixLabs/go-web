package model

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type EntityProperty interface {
	Name() string
	Type() string
	Value() any
	SetValue(name string, propType string, value *string)
	StringValue() *string
}

type entityProperty struct {
	state *entityPropertyState
}

type entityPropertyState struct {
	EntityId              uuid.UUID
	PropertyName          string
	PropertyType          string
	PropertyStringName    string
	PropertyIntName       string
	PropertyFloatName     string
	PropertyBoolName      string
	PropertyDateTimeName  string
	PropertyStringValue   string
	PropertyIntValue      uint64
	PropertyFloatValue    float64
	PropertyBoolValue     bool
	PropertyDateTimeValue time.Time
	Date                  time.Time
	IsNull                bool
	isDirty               bool
}

func (v *entityProperty) Name() string { return v.state.PropertyName }

func (v *entityProperty) Type() string {
	return v.state.PropertyType
}

func (v *entityProperty) Value() any {

	if v.state.IsNull {
		return nil
	}

	switch v.state.PropertyType {
	case "string":
		return v.state.PropertyStringValue
	case "int":
		return v.state.PropertyIntValue
	case "float":
		return v.state.PropertyFloatValue
	case "datetime":
		return v.state.PropertyDateTimeValue
	case "bool":
		return v.state.PropertyBoolValue
	}

	return nil
}

func (v *entityProperty) StringValue() *string {
	val := v.Value()

	if val != nil {
		strVal := fmt.Sprintf("%v", val)
		return &strVal
	}

	return nil
}

func (v *entityProperty) SetValue(name string, propType string, value *string) {

	comparedPropState := *v.state

	v.state.PropertyName = name
	v.state.PropertyStringName = ""
	v.state.PropertyIntName = ""
	v.state.PropertyFloatName = ""
	v.state.PropertyDateTimeName = ""
	v.state.PropertyBoolName = ""

	v.state.IsNull = true
	v.state.PropertyStringValue = ""
	v.state.PropertyBoolValue = false
	v.state.PropertyFloatValue = 0
	v.state.PropertyIntValue = 0
	v.state.PropertyDateTimeValue = time.UnixMilli(0).UTC()
	v.state.PropertyType = propType

	if value != nil {
		switch v.state.PropertyType {
		case "string":
			v.state.PropertyStringName = name
			v.state.PropertyStringValue = *value
		case "int":
			v.state.PropertyIntName = name
			intVal, err := strconv.ParseInt(*value, 10, 64)
			if err != nil {
				panic(err)
			}
			v.state.PropertyIntValue = convertToUint64(intVal)
		case "float":
			v.state.PropertyFloatName = name
			floatVal, err := strconv.ParseFloat(*value, 64)
			if err != nil {
				panic(err)
			}
			v.state.PropertyFloatValue = floatVal
		case "datetime":
			v.state.PropertyDateTimeName = name
			val, err := time.Parse(time.RFC3339, *value)
			if err != nil {
				panic(err)
			}
			v.state.PropertyDateTimeValue = val
		case "bool":
			v.state.PropertyBoolName = name
			boolVal, err := strconv.ParseBool(*value)
			if err != nil {
				panic(err)
			}
			v.state.PropertyBoolValue = boolVal
		default:
			panic("Type is not found:" + v.state.PropertyType)
		}

		v.state.IsNull = false
	}

	if comparedPropState != *v.state {
		v.state.Date = time.Now().UTC()
		v.state.isDirty = true
	}

}

func convertToUint64(v any) uint64 {
	switch n := v.(type) {
	case int:
		return uint64(n)
	case int8:
		return uint64(n)
	case int16:
		return uint64(n)
	case int32:
		return uint64(n)
	case int64:
		return uint64(n)
	case uint:
		return uint64(n)
	case uintptr:
		return uint64(n)
	case uint8:
		return uint64(n)
	case uint16:
		return uint64(n)
	case uint32:
		return uint64(n)
	case uint64:
		return n
	}

	panic("Can not convert to uint64")
}

type EntityPropertyRepository interface {
	Save(properties []EntityProperty)
	CreateBatch() driver.Batch
	SendBatch(batch driver.Batch)
	Add(batch driver.Batch, properties []EntityProperty)
	Find(whereAndPart []string, offset int64, totalRows int64, sortField string, sortFieldType string, sortOrder bool) EntityPropertyIterator
	FindAllUniqueNames() []string
}

type EntityPropertyIterator interface {
	Total() uint64
	Next() bool
	Get() EntityProperty
}

type EntityPropertyDto struct {
	Name     string  `json:"name"`
	Value    *string `json:"value"`
	Disabled bool    `json:"disabled"`
	Type     string  `json:"type"`
}

func GetEntityPropertyDto(entityProp EntityProperty) EntityPropertyDto {
	return EntityPropertyDto{
		Name:     entityProp.Name(),
		Value:    entityProp.StringValue(),
		Disabled: entityProp.Value() == nil,
		Type:     entityProp.Type(),
	}
}
