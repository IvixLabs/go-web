package model

import (
	"gorm.io/gorm"
)

func InitModels(db *gorm.DB) {
	AutoMigrateUser(db)
	AutoMigrateProduct(db)
}
