package usecase

import (
	"ivixlabs.com/goweb/internal/model/product"
)

//go:generate mockery --name ProductCreation
type ProductCreation interface {
	CreateNewProduct(form *product.Form, userId string) (product.Product, error)
}

type productCreationUseCase struct {
	productRepository product.Repository
}

func NewProductCreation(productRepository product.Repository) ProductCreation {
	return &productCreationUseCase{productRepository}
}

func (p *productCreationUseCase) CreateNewProduct(form *product.Form, userId string) (product.Product, error) {

	dto, err := form.GetUpdateProductDto()

	if err != nil {
		return nil, err
	}

	productObj := product.New(userId, form.Title, dto.Price, form.Brand)
	p.productRepository.CreateProduct(productObj)

	return productObj, nil
}
