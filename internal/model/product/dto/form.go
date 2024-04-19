package dto

import (
	"ivixlabs.com/goweb/internal/model/product"
)

type BaseForm struct {
	Title string `json:"title" validate:"required"`
	Brand string `json:"brand" validate:"required"`
	Price int    `json:"price" validate:"required,number"`
}

type CreateForm struct {
	UserId string `json:"userId" validate:"required,user_not_exists"`
	BaseForm
}

type UpdateForm struct {
	Id string `json:"id" validate:"required"`
	BaseForm
}

func (form *UpdateForm) GetUpdateDto() (*product.UpdateProductDto, error) {

	return &product.UpdateProductDto{
		Title:   form.Title,
		IsTitle: true,
		Brand:   form.Brand,
		IsBrand: true,
		Price:   form.Price,
		IsPrice: true,
	}, nil
}
