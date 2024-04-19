package product

import (
	"errors"
	"gorm.io/gorm"
	"ivixlabs.com/goweb/internal/model/product"
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

	repo.db.Preload("User").Find(&productArr)

	result := make([]product.Product, len(productArr))
	for i, productItem := range productArr {
		result[i] = product.FromState(productItem)
	}

	return result
}

func (repo *repository) FindProductById(productId string) product.Product {

	var productState product.State

	tx := repo.db.Preload("User").First(&productState, "id=?", productId)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return product.FromState(productState)
}

func (repo *repository) UpdateProduct(p product.Product) {
	repo.db.Omit("User").Save(p.State())
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

func (repo *repository) GetProductById(productId string) (product.Product, error) {
	var productState product.State

	tx := repo.db.Where("id=?", productId).First(&productState)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return product.FromState(productState), nil
}

func (repo *repository) DeleteProductById(productId string) {
	tx := repo.db.Delete(&product.State{}, "id=?", productId)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func (repo *repository) SaveProduct(p product.Product) {
	repo.db.Save(p.State())
}
