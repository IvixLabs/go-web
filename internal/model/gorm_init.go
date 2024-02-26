package model

import (
	"gorm.io/gorm"
)

func GormInitModels(db *gorm.DB) {
	autoMigrateUser(db)
	autoMigrateProduct(db)
}

func autoMigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&user{})
	if err != nil {
		panic(err)
	}
}

func autoMigrateProduct(db *gorm.DB) {
	err := db.AutoMigrate(&product{})
	if err != nil {
		panic(err)
	}
}
