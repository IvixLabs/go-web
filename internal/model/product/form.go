package product

import (
	"strconv"
)

type Form struct {
	Title string `validate:"required"`
	Brand string `validate:"required"`
	Price string `validate:"required,number"`
}

type PriceError struct {
	Value string
}

func (pr *PriceError) Error() string {
	return "price: wrong value = " + string(pr.Value)
}

func (form *Form) GetUpdateProductDto() (*UpdateProductDto, error) {

	price, err := strconv.Atoi(form.Price)
	if err != nil {
		return nil, &PriceError{form.Price}
	}

	return &UpdateProductDto{
		Title:   form.Title,
		IsTitle: true,
		Brand:   form.Brand,
		IsBrand: true,
		Price:   price,
		IsPrice: true,
	}, nil
}
