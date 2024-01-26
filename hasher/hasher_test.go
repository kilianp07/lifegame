package hasher

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	password := "password123"
	hash, err := Hash(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Hashed password does not match original password")
	}
}
