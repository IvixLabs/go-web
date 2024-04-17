package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ivixlabs.com/goweb/internal/model/product"
	mocks2 "ivixlabs.com/goweb/internal/model/product/mocks"
	"strconv"
	"testing"
)

func TestProductCreationCreate(t *testing.T) {
	repo := mocks2.NewRepository(t)
	repo.On("CreateProduct", mock.Anything).Once()

	useCase := NewProductCreation(repo)

	const userId = "123123"
	const price = 1000
	const title = "ProductTitle1"
	const brand = "ProductBrand1"

	form := &product.Form{
		Title: title,
		Brand: brand,
		Price: strconv.Itoa(price),
	}

	prodObj, err := useCase.CreateNewProduct(form, userId)

	assert.NoError(t, err)
	assert.NotEqual(t, prodObj.Id(), "")

	const wrongPrice = "aaabbb"

	form = &product.Form{
		Title: title,
		Brand: brand,
		Price: wrongPrice,
	}

	prodObj, err = useCase.CreateNewProduct(form, userId)
	assert.Equal(t, err, &product.PriceError{Value: wrongPrice})
}
