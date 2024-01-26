package model

import (
	"github.com/kilianp07/lifegame/hasher"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

// Hash hashes the user's password using a secure hashing algorithm.
// It returns an error if the hashing process fails.
func (u *User) Hash() error {
	hash, err := hasher.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}
