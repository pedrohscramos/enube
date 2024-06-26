package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Pedro Ramos", "teste@teste.com.br", "123456")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Pedro Ramos", user.Name)
	assert.Equal(t, "teste@teste.com.br", user.Email)
	assert.NotNil(t, user)
}

func TestUser_ValidadePassword(t *testing.T) {
	user, err := NewUser("Pedro Ramos", "teste@teste.com.br", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("7654321"))
	assert.NotEqual(t, "123456", user.Password)
}
