package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Product{}))

	product, err := entity.NewProduct("Product 1", "Description 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Description, productFound.Description)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Product{}))

	product, err := entity.NewProduct("Product 1", "Description 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Description, productFound.Description)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Product{}))
	productDB := NewProduct(db)

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), fmt.Sprintf("Description %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = productDB.Create(product)
		assert.NoError(t, err)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[2].Name, "Product 23")
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Product{}))

	product, err := entity.NewProduct("Product 1", "Description 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Product 1 Updated"
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Description, productFound.Description)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.Product{}))

	product, err := entity.NewProduct("Product 1", "Description 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
