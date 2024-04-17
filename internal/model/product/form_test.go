package product

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForm(t *testing.T) {
	const title = "Title1"
	const brand = "Brand1"
	const strPrice = "888"
	const intPrice = 888

	form := &Form{
		Title: title,
		Brand: brand,
		Price: strPrice,
	}

	dto, err := form.GetUpdateProductDto()
	assert.NoError(t, err)
	assert.Equal(t, dto.Title, title)
	assert.Equal(t, dto.Brand, brand)
	assert.Equal(t, dto.Price, intPrice)

	const wrongPrice = "aaabbb"
	const errorMsg = "price: wrong value = aaabbb"

	form = &Form{
		Title: title,
		Brand: brand,
		Price: wrongPrice,
	}

	dto, err = form.GetUpdateProductDto()
	assert.Nil(t, dto)
	assert.Equal(t, err, &PriceError{Value: wrongPrice})
	assert.Equal(t, err.Error(), errorMsg)
}
