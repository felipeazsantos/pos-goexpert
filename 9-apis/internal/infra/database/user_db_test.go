package database

import (
	"testing"

	"github.com/felipeazsantos/pos-goexpert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(&entity.User{}))
	user, _ := entity.NewUser("John Doe", "john@doe.com", "123456")

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.NoError(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NoError(t, db.AutoMigrate(entity.User{}))
	user, _ := entity.NewUser("John Doe", "john@doe.com", "123456")
	
	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.NoError(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.NotNil(t, userFound.Password)
	assert.NotEqual(t, "123456", userFound.Password)
}
