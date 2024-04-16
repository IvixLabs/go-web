package model

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type Entity interface {
	Id() uuid.UUID
	GetProperty(name string) any
	GetPropertyNames() []string
	SetProperty(name string, propType string, value *string)
}

type entity struct {
	state *entityState
}

type entityState struct {
	Id         uuid.UUID
	properties map[string]*entityProperty
}

func (e *entity) Id() uuid.UUID {
	return e.state.Id
}

func (e *entity) GetProperty(name string) any {

	if v, ok := e.state.properties[name]; ok {
		return v.Value()
	}

	return nil
}

func (e *entity) GetPropertyNames() []string {
	keys := make([]string, len(e.state.properties))

	i := 0
	for k := range e.state.properties {
		keys[i] = k
		i++
	}

	return keys
}

func (e *entity) SetProperty(name string, propType string, value *string) {

	v, ok := e.state.properties[name]
	if !ok {
		v = &entityProperty{&entityPropertyState{EntityId: e.state.Id}}
		e.state.properties[name] = v
	}

	v.SetValue(name, propType, value)
}

type EntityRepository interface {
	Save(entity Entity)
	CreateBatch() driver.Batch
	SendBatch(batch driver.Batch)
	Add(batch driver.Batch, entity Entity)
	Find(whereAndPart []string, offset int64, rows int64, sortField string, sortFieldType string, sortOrder bool) ([]Entity, uint64)
}

func GetEntityDto(en Entity) EntityDto {
	dto := EntityDto{Id: en.Id().String(), Properties: make(map[string]EntityPropertyDto)}

	entityStruct := en.(*entity)

	for _, prop := range entityStruct.state.properties {
		dto.Properties[prop.Name()] = GetEntityPropertyDto(prop)
	}

	return dto
}

type EntityDto struct {
	Id         string                       `json:"id"`
	Properties map[string]EntityPropertyDto `json:"properties"`
}

func NewEntity() Entity {
	return &entity{
		state: &entityState{
			Id:         uuid.New(),
			properties: make(map[string]*entityProperty),
		},
	}
}
