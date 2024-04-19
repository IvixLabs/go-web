package product

import (
	"ivixlabs.com/goweb/internal/model/product"
)

type Dto struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Brand     string `json:"brand"`
	Price     int    `json:"price"`
	UserEmail string `json:"userEmail"`
}

func GetDto(productObj product.Product) *Dto {
	return &Dto{
		Id:        productObj.Id(),
		Title:     productObj.Title(),
		Brand:     productObj.Brand(),
		Price:     productObj.Price(),
		UserEmail: productObj.User().Email(),
	}
}
