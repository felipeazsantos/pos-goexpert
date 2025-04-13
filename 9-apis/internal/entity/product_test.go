package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", "Description 1", 10)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, "Description 1", product.Description)
	assert.Equal(t, 10, product.Price)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("", "Description 1", 10)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired.Error(), err.Error())
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	product, err := NewProduct("", "Description 1", 10)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired.Error(), err.Error())
}

func TestProductValidateInvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 1", "Description 1", -1)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice.Error(), err.Error())
}

func TestProductWhenPriceIsZero(t *testing.T) {
	product, err := NewProduct("Product 1", "Description 1", 0)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired.Error(), err.Error())
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	product, err := NewProduct("Product 1", "Description 1", -1)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice.Error(), err.Error())
}



