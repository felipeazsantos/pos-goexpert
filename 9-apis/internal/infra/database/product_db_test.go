package database

import (
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

	product1, err := entity.NewProduct("Product 1", "Description 1", 10)
	assert.NoError(t, err)

	product2, err := entity.NewProduct("Product 2", "Description 2", 20)
	assert.NoError(t, err)

	product3, err := entity.NewProduct("Product 3", "Description 3", 30)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product1)
	assert.NoError(t, err)
	err = productDB.Create(product2)
	assert.NoError(t, err)
	err = productDB.Create(product3)
	assert.NoError(t, err)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, product1.ID, products[0].ID)
	assert.Equal(t, product2.ID, products[1].ID)
	assert.Equal(t, product3.ID, products[2].ID)
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
