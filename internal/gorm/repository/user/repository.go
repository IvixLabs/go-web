package user

import (
	"errors"
	"ivixlabs.com/goweb/internal/model/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &repository{db: db}
}

func (repo *repository) SaveUser(u user.User) {
	repo.db.Save(u.State())
}

func (repo *repository) FindAllUsers() []user.User {

	var userArr []user.State

	repo.db.Find(&userArr)

	result := make([]user.User, len(userArr))
	for i, userItem := range userArr {
		result[i] = user.FromState(userItem)
	}

	return result
}

func (repo *repository) FindUserByEmail(email string) user.User {

	var userObj user.State

	tx := repo.db.Where("email=?", email).First(&userObj)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return user.FromState(userObj)
}

func (repo *repository) GetUserById(userId string) (user.User, error) {
	var userObj user.State

	tx := repo.db.Where("id=?", userId).First(&userObj)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return user.FromState(userObj), nil
}

func (repo *repository) DeleteUserById(userId string) {
	tx := repo.db.Delete(&user.State{}, "id=?", userId)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
