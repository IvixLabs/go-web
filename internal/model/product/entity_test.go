package product

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	const userId = "111"
	const title = "Title1"
	const price = 999
	const brand = "Brand1"

	prod := New(userId, title, price, brand)

	assert.Equal(t, prod.State().TableName(), "product")
	assert.NotEqual(t, prod.Id(), "")
	assert.Equal(t, prod.UserId(), userId)
	assert.Equal(t, prod.Title(), title)
	assert.Equal(t, prod.Price(), price)
	assert.Equal(t, prod.Brand(), brand)

	const newTitle = "Title1"
	const newPrice = 999
	const newBrand = "Brand1"

	prod.Update(&UpdateProductDto{
		Title:   newTitle,
		IsTitle: false,
		Price:   newPrice,
		IsPrice: false,
		Brand:   newBrand,
		IsBrand: false,
	})

	assert.Equal(t, prod.Title(), title)
	assert.Equal(t, prod.Price(), price)
	assert.Equal(t, prod.Brand(), brand)

	prod.Update(&UpdateProductDto{
		Title:   newTitle,
		IsTitle: true,
		Price:   newPrice,
		IsPrice: true,
		Brand:   newBrand,
		IsBrand: true,
	})

	assert.Equal(t, prod.Title(), title)
	assert.Equal(t, prod.Price(), price)
	assert.Equal(t, prod.Brand(), brand)

}
