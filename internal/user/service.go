package user

import (
	"gorm.io/gorm"
	"ivixlabs.com/goweb/internal/model"
	"ivixlabs.com/goweb/internal/user/registration"
)

type Service interface {
	CreateNewUser(form *registration.Form) (model.User, error)
	FindAll() []model.User
	FindByEmail(email string) model.User
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (service *service) FindAll() []model.User {
	return model.FindAllUsers(service.db)
}

func (service *service) FindByEmail(email string) model.User {
	return model.FindUserByEmail(service.db, email)
}

func (service *service) CreateNewUser(form *registration.Form) (model.User, error) {
	userObj := model.NewUser(form.Email, form.Password, form.Address)
	model.SaveUser(userObj, service.db)

	return userObj, nil
}
