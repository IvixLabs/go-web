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
	repo.db.Create(pStruct.state)
}

func (repo *gormProductRepository) FindAllProducts() []Product {

	var productArr []productState

	repo.db.Find(&productArr)

	result := make([]Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = &product{state: productItem}
	}

	return result
}

func (repo *gormProductRepository) FindProductById(productId string) Product {

	var productState productState

	tx := repo.db.First(&productState, "id=?", productId)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &product{state: productState}
}

func (repo *gormProductRepository) UpdateProduct(p Product, updateProduct *UpdateProductArg) {
	pStruct := (p).(*product)

	if updateProduct.IsTitle {
		pStruct.state.Title = updateProduct.Title
	}

	if updateProduct.IsBrand {
		pStruct.state.Brand = updateProduct.Brand
	}

	if updateProduct.IsPrice {
		pStruct.state.Price = updateProduct.Price
	}

	repo.db.Save(pStruct.state)
}

func (repo *gormProductRepository) DeleteProduct(id string, userId string) int64 {
	tx := repo.db.Delete(&productState{}, "id=? and user_id=?", id, userId)
	return tx.RowsAffected
}

func (repo *gormProductRepository) FindProductsByUserId(userId string) []Product {

	var productArr []productState

	repo.db.Find(&productArr, "user_id=?", userId)

	result := make([]Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = &product{state: productItem}
	}

	return result
}
