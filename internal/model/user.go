package model

import (
	"github.com/google/uuid"
)

type User interface {
	GetId() string
	GetEmail() string
	GetAddress() string
	GetPassword() string
}

type UserRepository interface {
	SaveUser(u User)
	FindAllUsers() []User
	FindUserByEmail(email string) User
}

type user struct {
	Id       string `gorm:"primaryKey"`
	Email    string
	Address  string
	Password string
}

func (u *user) GetId() string {
	return u.Id
}

func (u *user) GetEmail() string {
	return u.Email
}

func (u *user) GetAddress() string {
	return u.Address
}

func (u *user) GetPassword() string { return u.Password }

func NewUser(email string, password string, address string) User {

	return &user{
		Id:       uuid.NewString(),
		Email:    email,
		Password: password,
		Address:  address,
	}
}
