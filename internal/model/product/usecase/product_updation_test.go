package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ivixlabs.com/goweb/internal/model/product"
	mocks2 "ivixlabs.com/goweb/internal/model/product/mocks"
	"strconv"
	"testing"
)

func TestProductUpdatingUpdate(t *testing.T) {
	repo := mocks2.NewRepository(t)
	repo.On("UpdateProduct", mock.Anything, mock.Anything).Once()

	useCase := NewProductUpdating(repo)

	const price = 2000
	const title = "Title2"
	const brand = "Brand2"
	const userId = "123"

	form := &product.Form{
		Title: title,
		Brand: brand,
		Price: strconv.Itoa(price),
	}

	productObj := product.New(userId, "Title1", 1000, "Brand1")

	err := useCase.UpdateProduct(form, productObj)

	assert.NoError(t, err)

	const wrongPrice = "abc"
	form = &product.Form{
		Title: title,
		Brand: brand,
		Price: wrongPrice,
	}

	err = useCase.UpdateProduct(form, productObj)
	assert.Equal(t, err, &product.PriceError{Value: wrongPrice})

}
