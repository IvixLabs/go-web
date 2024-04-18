package dto

import (
	"ivixlabs.com/goweb/internal/model/user"
)

type CreateForm struct {
	Email    string `json:"email" validate:"required,user_email_exists"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
}

type UpdateForm struct {
	Id       string `json:"id" validate:"required"`
	Password string `json:"password"`
	Address  string `json:"address" validate:"required"`
}

func (form *UpdateForm) GetUpdateDto() (*user.UpdateDto, error) {

	return &user.UpdateDto{
		Address:    form.Address,
		IsAddress:  true,
		Password:   form.Password,
		IsPassword: form.Password != "",
	}, nil
}
