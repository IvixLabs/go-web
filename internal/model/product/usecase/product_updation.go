package usecase

import (
	"ivixlabs.com/goweb/internal/model/product"
)

type ProductUpdating interface {
	UpdateProduct(form *product.Form, p product.Product) error
}

type productUpdatingUseCase struct {
	productRepository product.Repository
}

func NewProductUpdating(productRepository product.Repository) ProductUpdating {
	return &productUpdatingUseCase{productRepository}
}

func (pu *productUpdatingUseCase) UpdateProduct(form *product.Form, p product.Product) error {

	updateP, err := form.GetUpdateProductDto()

	if err != nil {
		return err
	}

	p.Update(updateP)

	pu.productRepository.UpdateProduct(p)

	return nil
}
