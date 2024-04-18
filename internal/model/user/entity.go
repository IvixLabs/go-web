package user

import (
	"github.com/google/uuid"
	"ivixlabs.com/goweb/internal/model"
)

type User interface {
	model.Model[State]
	Id() string
	Email() string
	Address() string
	Password() string
	Update(updateDto *UpdateDto)
}

type State struct {
	Id       string `gorm:"primaryKey"`
	Email    string
	Address  string
	Password string
}

func (State) TableName() string {
	return "user"
}

type user struct {
	model.BaseModel[State]
}

func (u *user) Id() string {
	return u.State().Id
}

func (u *user) Email() string {
	return u.State().Email
}

func (u *user) Address() string {
	return u.State().Address
}

func (u *user) Password() string {
	return u.State().Password
}

func (u *user) Update(updateDto *UpdateDto) {
	state := u.State()

	if updateDto.IsAddress {
		state.Address = updateDto.Address
	}

	if updateDto.IsPassword {
		state.Password = updateDto.Password
	}

	u.UpdateState(state)
}

func FromState(state State) User {
	u := &user{}
	u.UpdateState(state)
	return u
}

func New(email string, password string, address string) User {

	return FromState(State{
		Id:       uuid.NewString(),
		Email:    email,
		Password: password,
		Address:  address,
	})

}

type UpdateDto struct {
	Address    string
	IsAddress  bool
	Password   string
	IsPassword bool
}
