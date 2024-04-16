package property

import (
	"github.com/google/uuid"
	"ivixlabs.com/goweb/internal/model"
	"ivixlabs.com/goweb/internal/model/user"
)

type Property interface {
	model.Model[State]
	Id() string
	Name() string
}

type property struct {
	model.BaseModel[State]
}

type State struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	UserId string
	User   user.State `gorm:"references:Id"`
}

type PropertyDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetPropertyDto(p Property) PropertyDto {
	state := p.State()
	return PropertyDto{
		Id:   state.Id,
		Name: state.Name,
	}
}

func (State) TableName() string {
	return "property"
}

func (p *property) Id() string {
	return p.State().Id
}

func (p *property) Name() string {
	return p.State().Name
}

func (p *property) UserId() string {
	return p.State().UserId
}

func FromState(state State) Property {
	p := &property{}
	p.UpdateState(state)
	return p
}

func New(userId string, name string) Property {

	return FromState(State{
		Id:     uuid.NewString(),
		UserId: userId,
		Name:   name,
	})

}
