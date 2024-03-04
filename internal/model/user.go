package model

import (
	"github.com/google/uuid"
)

type User interface {
	Id() string
	Email() string
	Address() string
	Password() string
}

type UserRepository interface {
	SaveUser(u User)
	FindAllUsers() []User
	FindUserByEmail(email string) User
}

type userState struct {
	Id       string `gorm:"primaryKey"`
	Email    string
	Address  string
	Password string
}

func (userState) TableName() string {
	return "user"
}

type user struct {
	state userState
}

func (u *user) Id() string {
	return u.state.Id
}

func (u *user) Email() string {
	return u.state.Email
}

func (u *user) Address() string {
	return u.state.Address
}

func (u *user) Password() string {
	return u.state.Password
}

func NewUser(email string, password string, address string) User {

	return &user{
		state: userState{
			Id:       uuid.NewString(),
			Email:    email,
			Password: password,
			Address:  address,
		},
	}
}
