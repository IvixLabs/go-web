package product

import (
	"errors"
	"ivixlabs.com/goweb/internal/model/product"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.Repository {
	return &repository{db: db}
}

func (repo *repository) CreateProduct(p product.Product) {
	repo.db.Create(p.State())
}

func (repo *repository) FindAllProducts() []product.Product {

	var productArr []product.State

	repo.db.Find(&productArr)

	result := make([]product.Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = product.FromState(productItem)
	}

	return result
}

func (repo *repository) FindProductById(productId string) product.Product {

	var productState product.State

	tx := repo.db.First(&productState, "id=?", productId)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return product.FromState(productState)
}

func (repo *repository) UpdateProduct(p product.Product) {
	repo.db.Save(p.State())
}

func (repo *repository) DeleteProduct(id string, userId string) int64 {
	tx := repo.db.Delete(&product.State{}, "id=? and user_id=?", id, userId)
	return tx.RowsAffected
}

func (repo *repository) FindProductsByUserId(userId string) []product.Product {

	var productArr []product.State

	repo.db.Find(&productArr, "user_id=?", userId)

	result := make([]product.Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = product.FromState(productItem)
	}

	return result
}
