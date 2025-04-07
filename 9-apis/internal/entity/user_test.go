package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "123456")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jane Doe", "jane@email.com", "654321")
	assert.NoError(t, err)
	assert.True(t, user.ValidatePassword("654321"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, "654321", user.Password)
}
