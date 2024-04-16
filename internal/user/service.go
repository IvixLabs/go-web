package user

import (
	"ivixlabs.com/goweb/internal/model/user"
	"ivixlabs.com/goweb/internal/user/registration"
)

type Service interface {
	CreateNewUser(form *registration.Form) (user.User, error)
	FindAll() []user.User
	FindByEmail(email string) user.User
}

type service struct {
	userRepository user.Repository
}

func NewService(userRepository user.Repository) Service {
	return &service{userRepository: userRepository}
}

func (service *service) FindAll() []user.User {
	return service.userRepository.FindAllUsers()
}

func (service *service) FindByEmail(email string) user.User {
	return service.userRepository.FindUserByEmail(email)
}

func (service *service) CreateNewUser(form *registration.Form) (user.User, error) {
	userObj := user.New(form.Email, form.Password, form.Address)
	service.userRepository.SaveUser(userObj)
	return userObj, nil
}
