package model

import (
	"errors"

	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (repo *gormUserRepository) SaveUser(u User) {
	uStruct := (u).(*user)
	repo.db.Create(uStruct.state)
}

func (repo *gormUserRepository) FindAllUsers() []User {

	var userArr []userState

	repo.db.Find(&userArr)

	result := make([]User, len(userArr))
	for i, userItem := range userArr {
		result[i] = &user{state: userItem}
	}

	return result
}

func (repo *gormUserRepository) FindUserByEmail(email string) User {

	var userObj userState

	tx := repo.db.Where("email=?", email).First(&userObj)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user{state: userObj}
}
