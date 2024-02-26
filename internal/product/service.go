package product

import (
	"strconv"

	"ivixlabs.com/goweb/internal/model"
)

type Service interface {
	CreateNewProduct(form *Form, userId string) model.Product
	FindAll() []model.Product
	FindByUserId(userId string) []model.Product
	FindById(productId string) model.Product
	UpdateProduct(form *Form, p model.Product)
	DeleteProduct(id string, userId string) int64
}

type service struct {
	productRepository model.ProductRepository
}

func NewService(productRepository model.ProductRepository) Service {
	return &service{productRepository: productRepository}
}

func (service *service) CreateNewProduct(form *Form, userId string) model.Product {

	price, err := strconv.Atoi(form.Price)
	if err != nil {
		panic(err)
	}

	productObj := model.NewProduct(userId, form.Title, price, form.Brand)
	service.productRepository.Create(productObj)

	return productObj
}

func (service *service) UpdateProduct(form *Form, p model.Product) {
	price, err := strconv.Atoi(form.Price)
	if err != nil {
		panic(err)
	}

	updateP := model.UpdateProductArg{
		Title:   form.Title,
		IsTitle: true,
		Brand:   form.Brand,
		IsBrand: true,
		Price:   price,
		IsPrice: true,
	}

	service.productRepository.UpdateProduct(p, &updateP)
}

func (service *service) FindAll() []model.Product {
	return service.productRepository.FindAllProducts()
}

func (service *service) FindByUserId(userId string) []model.Product {
	return service.productRepository.FindProductsByUserId(userId)
}

func (service *service) FindById(productId string) model.Product {
	return service.productRepository.FindProductById(productId)
}

func (service *service) DeleteProduct(id string, userId string) int64 {
	return service.productRepository.DeleteProduct(id, userId)
}
