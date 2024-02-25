package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product interface {
	GetId() string
	GetTitle() string
	GetPrice() int
	GetBrand() string
	GetUserId() string
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

func Create(p Product, db *gorm.DB) {
	pStruct := (p).(*product)
	db.Create(pStruct)
}

func FindAllProducts(db *gorm.DB) []Product {

	var productArr []product

	db.Find(&productArr)

	result := make([]Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = &productItem
	}

	return result
}

func FindProductById(db *gorm.DB, productId string) Product {

	var productObj product

	tx := db.First(&productObj, "id=?", productId)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &productObj
}

func AutoMigrateProduct(db *gorm.DB) {
	err := db.AutoMigrate(&product{})
	if err != nil {
		panic(err)
	}
}

func UpdateProduct(p Product, updateProduct *UpdateProductArg, db *gorm.DB) {
	pStruct := (p).(*product)

	if updateProduct.IsTitle {
		pStruct.Title = updateProduct.Title
	}

	if updateProduct.IsBrand {
		pStruct.Brand = updateProduct.Brand
	}

	if updateProduct.IsPrice {
		pStruct.Price = updateProduct.Price
	}

	db.Save(p)
}

func DeleteProduct(id string, userId string, db *gorm.DB) int64 {
	tx := db.Delete(&product{}, "id=? and user_id=?", id, userId)
	return tx.RowsAffected
}

func FindProductsByUserId(db *gorm.DB, userId string) []Product {

	var productArr []product

	db.Find(&productArr, "user_id=?", userId)

	result := make([]Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = &productItem
	}

	return result
}
