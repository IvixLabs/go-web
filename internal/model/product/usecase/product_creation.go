package usecase

import (
	"ivixlabs.com/goweb/internal/model/product"
	"strconv"
)

type ProductCreation interface {
	CreateNewProduct(form *product.Form, userId string) product.Product
}

type productCreationUseCase struct {
	productRepository product.Repository
}

func NewProductCreation(productRepository product.Repository) ProductCreation {
	return &productCreationUseCase{productRepository}
}

func (p *productCreationUseCase) CreateNewProduct(form *product.Form, userId string) product.Product {

	price, err := strconv.Atoi(form.Price)
	if err != nil {
		panic(err)
	}

	productObj := product.New(userId, form.Title, price, form.Brand)
	p.productRepository.Create(productObj)

	return productObj
}
