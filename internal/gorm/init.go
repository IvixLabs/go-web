package gorm

import (
	"gorm.io/gorm"
	"ivixlabs.com/goweb/internal/model/product"
	"ivixlabs.com/goweb/internal/model/property"
	"ivixlabs.com/goweb/internal/model/user"
)

func InitModels(db *gorm.DB) {
	autoMigrateUser(db)
	autoMigrateProduct(db)
	autoMigrateProperty(db)
}

func autoMigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&user.State{})
	if err != nil {
		panic(err)
	}
}

func autoMigrateProduct(db *gorm.DB) {
	err := db.AutoMigrate(&product.State{})
	if err != nil {
		panic(err)
	}
}

func autoMigrateProperty(db *gorm.DB) {
	err := db.AutoMigrate(&property.State{})
	if err != nil {
		panic(err)
	}
}
