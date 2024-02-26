package user

import (
	"ivixlabs.com/goweb/internal/model"
	"ivixlabs.com/goweb/internal/user/registration"
)

type Service interface {
	CreateNewUser(form *registration.Form) (model.User, error)
	FindAll() []model.User
	FindByEmail(email string) model.User
}

type service struct {
	userRepository model.UserRepository
}

func NewService(userRepository model.UserRepository) Service {
	return &service{userRepository: userRepository}
}

func (service *service) FindAll() []model.User {
	return service.userRepository.FindAllUsers()
}

func (service *service) FindByEmail(email string) model.User {
	return service.userRepository.FindUserByEmail(email)
}

func (service *service) CreateNewUser(form *registration.Form) (model.User, error) {
	userObj := model.NewUser(form.Email, form.Password, form.Address)
	service.userRepository.SaveUser(userObj)
	return userObj, nil
}
