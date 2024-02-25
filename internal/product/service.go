package product

import (
	"strconv"

	"gorm.io/gorm"
	"ivixlabs.com/proj5/internal/model"
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
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (service *service) CreateNewProduct(form *Form, userId string) model.Product {

	price, err := strconv.Atoi(form.Price)
	if err != nil {
		panic(err)
	}

	productObj := model.NewProduct(userId, form.Title, price, form.Brand)
	model.Create(productObj, service.db)

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

	model.UpdateProduct(p, &updateP, service.db)
}

func (service *service) FindAll() []model.Product {
	return model.FindAllProducts(service.db)
}

func (service *service) FindByUserId(userId string) []model.Product {
	return model.FindProductsByUserId(service.db, userId)
}

func (service *service) FindById(productId string) model.Product {
	return model.FindProductById(service.db, productId)
}

func (service *service) DeleteProduct(id string, userId string) int64 {
	return model.DeleteProduct(id, userId, service.db)
}
