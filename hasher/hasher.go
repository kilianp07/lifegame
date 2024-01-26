package hasher

import "golang.org/x/crypto/bcrypt"

// Hash is a function that takes a password string and returns its hashed representation.
// It uses bcrypt.GenerateFromPassword to generate a secure hash with a cost factor of 14.
// The hashed password is returned as a string.
// If an error occurs during the hashing process, it is also returned.
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with its corresponding hash and returns true if they match.
// It uses bcrypt.CompareHashAndPassword to perform the comparison.
// The password and hash parameters should be strings.
// It returns a boolean value indicating whether the password matches the hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
