package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10, product.Price)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	product, err := NewProduct("", 10)
	fmt.Println(err.Error(), ErrNameIsRequired.Error())
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired.Error(), err.Error())
}

func TestProductValidateInvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestProductWhenPriceIsZero(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired.Error(), err.Error())
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice.Error(), err.Error())
}



