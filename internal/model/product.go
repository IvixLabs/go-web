package model

import (
	"github.com/google/uuid"
)

type Product interface {
	GetId() string
	GetTitle() string
	GetPrice() int
	GetBrand() string
	GetUserId() string
}

type ProductRepository interface {
	Create(p Product)
	FindAllProducts() []Product
	FindProductById(productId string) Product
	UpdateProduct(p Product, updateProduct *UpdateProductArg)
	DeleteProduct(id string, userId string) int64
	FindProductsByUserId(userId string) []Product
}

type UpdateProductArg struct {
	Title   string
	IsTitle bool
	Price   int
	IsPrice bool
	Brand   string
	IsBrand bool
}

type product struct {
	Id     string `gorm:"primaryKey"`
	Title  string
	Price  int
	Brand  string
	Info   string
	UserId string
	User   user `gorm:"references:Id"`
}

func (p *product) GetId() string {
	return p.Id
}

func (p *product) GetTitle() string {
	return p.Title
}

func (p *product) GetPrice() int {
	return p.Price
}

func (p *product) GetBrand() string {
	return p.Brand
}

func (p *product) GetUserId() string {
	return p.UserId
}

func NewProduct(userId string, title string, price int, brand string) Product {

	return &product{
		Id:     uuid.NewString(),
		UserId: userId,
		Title:  title,
		Price:  price,
		Brand:  brand,
		Info:   title + " " + brand,
	}
}
