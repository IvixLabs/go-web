package model

import (
	"github.com/google/uuid"
)

type Product interface {
	Id() string
	Title() string
	Price() int
	Brand() string
	UserId() string
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
	state productState
}

type productState struct {
	Id     string `gorm:"primaryKey"`
	Title  string
	Price  int
	Brand  string
	Info   string
	UserId string
	User   userState `gorm:"references:Id"`
}

func (productState) TableName() string {
	return "product"
}

func (p *product) Id() string {
	return p.state.Id
}

func (p *product) Title() string {
	return p.state.Title
}

func (p *product) Price() int {
	return p.state.Price
}

func (p *product) Brand() string {
	return p.state.Brand
}

func (p *product) UserId() string {
	return p.state.UserId
}

func NewProduct(userId string, title string, price int, brand string) Product {

	return &product{
		state: productState{
			Id:     uuid.NewString(),
			UserId: userId,
			Title:  title,
			Price:  price,
			Brand:  brand,
			Info:   title + " " + brand,
		},
	}
}
