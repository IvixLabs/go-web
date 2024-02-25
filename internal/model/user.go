package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	GetId() string
	GetEmail() string
	GetAddress() string
	GetPassword() string
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

func SaveUser(u User, db *gorm.DB) {
	pStruct := (u).(*user)
	db.Create(pStruct)
}

func FindAllUsers(db *gorm.DB) []User {

	var userArr []user

	db.Find(&userArr)

	result := make([]User, len(userArr))
	for i, userItem := range userArr {
		result[i] = &userItem
	}

	return result
}

func FindUserByEmail(db *gorm.DB, email string) User {

	var userObj user

	tx := db.Where("email=?", email).First(&userObj)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &userObj
}

func AutoMigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&user{})
	if err != nil {
		panic(err)
	}
}
