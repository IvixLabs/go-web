package property

import (
	"errors"
	"gorm.io/gorm"
	"ivixlabs.com/goweb/internal/model/property"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) property.Repository {
	return &repository{db: db}
}

func (repo *repository) Create(p property.Property) {
	repo.db.Create(p.State())
}

func (repo *repository) FindAll() []property.Property {

	var pArr []property.State

	repo.db.Find(&pArr)

	result := make([]property.Property, len(pArr))
	for i, productItem := range pArr {
		result[i] = property.FromState(productItem)
	}

	return result
}

func (repo *repository) FindById(id string) property.Property {

	var state property.State

	tx := repo.db.First(&state, "id=?", id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return property.FromState(state)
}

func (repo *repository) Update(p property.Property, updateP *property.UpdateProperty) {
	state := p.State()

	if updateP.IsName {
		state.Name = updateP.Name
	}

	repo.db.Save(state)
}

func (repo *repository) Delete(id string, userId string) int64 {
	tx := repo.db.Delete(&property.State{}, "id=? and user_id=?", id, userId)
	return tx.RowsAffected
}

func (repo *repository) FindByUserId(userId string) []property.Property {

	var arr []property.State

	repo.db.Find(&arr, "user_id=?", userId)

	result := make([]property.Property, len(arr))
	for i, item := range arr {
		result[i] = property.FromState(item)
	}

	return result
}
