package usecase

import (
	"ivixlabs.com/goweb/internal/model/product"
	"strconv"
)

type ProductUpdating interface {
	UpdateProduct(form *product.Form, p product.Product)
}

type productUpdatingUseCase struct {
	productRepository product.Repository
}

func NewProductUpdating(productRepository product.Repository) ProductUpdating {
	return &productUpdatingUseCase{productRepository}
}

func (pu *productUpdatingUseCase) UpdateProduct(form *product.Form, p product.Product) {

	price, err := strconv.Atoi(form.Price)
	if err != nil {
		panic(err)
	}

	updateP := product.UpdateProductArg{
		Title:   form.Title,
		IsTitle: true,
		Brand:   form.Brand,
		IsBrand: true,
		Price:   price,
		IsPrice: true,
	}

	pu.productRepository.UpdateProduct(p, &updateP)
}
