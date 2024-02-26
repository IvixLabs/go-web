package model

import (
	"errors"

	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ProductRepository {
	return &gormProductRepository{db: db}
}

func (repo *gormProductRepository) Create(p Product) {
	pStruct := (p).(*product)
	repo.db.Create(pStruct)
}

func (repo *gormProductRepository) FindAllProducts() []Product {

	var productArr []product

	repo.db.Find(&productArr)

	result := make([]Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = &productItem
	}

	return result
}

func (repo *gormProductRepository) FindProductById(productId string) Product {

	var productObj product

	tx := repo.db.First(&productObj, "id=?", productId)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &productObj
}

func (repo *gormProductRepository) UpdateProduct(p Product, updateProduct *UpdateProductArg) {
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

	repo.db.Save(p)
}

func (repo *gormProductRepository) DeleteProduct(id string, userId string) int64 {
	tx := repo.db.Delete(&product{}, "id=? and user_id=?", id, userId)
	return tx.RowsAffected
}

func (repo *gormProductRepository) FindProductsByUserId(userId string) []Product {

	var productArr []product

	repo.db.Find(&productArr, "user_id=?", userId)

	result := make([]Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = &productItem
	}

	return result
}
