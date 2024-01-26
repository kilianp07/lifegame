package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHash(t *testing.T) {
	const password = "password123"

	user := &User{
		Password: password,
	}

	err := user.Hash()
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Password)

	if password == user.Password {
		t.Errorf("Hashed password is not different from original password")
	}
}
