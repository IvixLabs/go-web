package model

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type clickHouseEntityRepository struct {
	conn                     driver.Conn
	entityPropertyRepository EntityPropertyRepository
}

func NewClickHouseEntityRepository(conn driver.Conn, entityPropertyRepository EntityPropertyRepository) EntityRepository {
	return &clickHouseEntityRepository{conn: conn, entityPropertyRepository: entityPropertyRepository}
}

func (repo *clickHouseEntityRepository) CreateBatch() driver.Batch {
	return repo.entityPropertyRepository.CreateBatch()
}

func (repo *clickHouseEntityRepository) SendBatch(batch driver.Batch) {
	repo.entityPropertyRepository.SendBatch(batch)
}

func (repo *clickHouseEntityRepository) Add(batch driver.Batch, entityObj Entity) {

	entityStruct := entityObj.(*entity)

	props := make([]EntityProperty, len(entityStruct.state.properties))
	i := 0
	for _, prop := range entityStruct.state.properties {
		props[i] = prop
		i++
	}

	repo.entityPropertyRepository.Add(batch, props)
}

func (repo *clickHouseEntityRepository) Save(entityObj Entity) {

	batch := repo.CreateBatch()
	repo.Add(batch, entityObj)
	repo.SendBatch(batch)
}

func (repo *clickHouseEntityRepository) Find(whereAndPart []string, offset int64, rows int64, sortField string, sortFieldType string, sortOrder bool) ([]Entity, uint64) {

	propertyIterator := repo.entityPropertyRepository.Find(whereAndPart, offset, rows, sortField, sortFieldType, sortOrder)

	result := make(map[uuid.UUID]*entity)
	sortedIds := make([]uuid.UUID, 0, 10)

	for propertyIterator.Next() {
		prop := propertyIterator.Get().(*entityProperty)

		if _, ok := result[prop.state.EntityId]; !ok {
			result[prop.state.EntityId] = &entity{
				state: &entityState{Id: prop.state.EntityId, properties: make(map[string]*entityProperty)},
			}
			sortedIds = append(sortedIds, prop.state.EntityId)
		}

		result[prop.state.EntityId].state.properties[prop.Name()] = prop
	}

	entities := make([]Entity, len(result))

	i := 0
	for _, v := range sortedIds {
		entities[i] = result[v]
		i++
	}

	return entities, propertyIterator.Total()
}
