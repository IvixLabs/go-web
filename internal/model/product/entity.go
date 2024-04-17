package product

import (
	"github.com/google/uuid"
	"ivixlabs.com/goweb/internal/model"
	"ivixlabs.com/goweb/internal/model/user"
)

type Product interface {
	model.Model[State]
	Id() string
	Title() string
	Price() int
	Brand() string
	UserId() string
	Update(updateProduct *UpdateProductDto)
}

type product struct {
	model.BaseModel[State]
}

type State struct {
	Id     string `gorm:"primaryKey"`
	Title  string
	Price  int
	Brand  string
	Info   string
	UserId string
	User   user.State `gorm:"references:Id"`
}

func (State) TableName() string {
	return "product"
}

func (p *product) Id() string {
	return p.State().Id
}

func (p *product) Title() string {
	return p.State().Title
}

func (p *product) Price() int {
	return p.State().Price
}

func (p *product) Brand() string {
	return p.State().Brand
}

func (p *product) UserId() string {
	return p.State().UserId
}

func (p *product) Update(updateProduct *UpdateProductDto) {
	state := p.State()

	if updateProduct.IsTitle {
		state.Title = updateProduct.Title
	}

	if updateProduct.IsBrand {
		state.Brand = updateProduct.Brand
	}

	if updateProduct.IsPrice {
		state.Price = updateProduct.Price
	}

	p.UpdateState(state)
}

func FromState(state State) Product {
	p := &product{}
	p.UpdateState(state)
	return p
}

func New(userId string, title string, price int, brand string) Product {

	return FromState(State{
		Id:     uuid.NewString(),
		UserId: userId,
		Title:  title,
		Price:  price,
		Brand:  brand,
		Info:   title + " " + brand,
	})
}

type UpdateProductDto struct {
	Title   string
	IsTitle bool
	Price   int
	IsPrice bool
	Brand   string
	IsBrand bool
}
