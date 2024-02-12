 # hasher Package Documentation

## Introduction
The `hasher` package is a simple Go implementation using the `bcrypt` library for securely hashing and comparing passwords.

## Dependencies
This package depends on the `golang.org/x/crypto/bcrypt` library for hashing functionality. Make sure to have this dependency installed in your project before using this package.

## Functions
### Hash(password string) ([]byte, error)
The `Hash` function takes a plaintext password as an argument and returns the corresponding salted and hashed byte array along with any potential error. This hash can be securely stored for later comparison when authenticating users.

```go
import "yourproject/hasher" // Replace 'yourproject' with your project name

password := "password123"
hash, err := hasher.Hash(password)
if err != nil {
	log.Println("Error hashing password:", err)
}
// Store the hash securely for later use
```

### bcrypt.CompareHashAndPassword(hashedPassword []byte, plaintextPassword []byte) error
The `bcrypt.CompareHashAndPassword` function from the `bcrypt` library is used to compare a given plaintext password against a previously hashed password (salt and hash). The function returns an error if there's a mismatch or if any other issue occurs during comparison.

```go
import "yourproject/hasher" // Replace 'yourproject' with your project name
import "golang.org/x/crypto/bcrypt"

password := "password123"
hash, err := hasher.Hash(password)
if err != nil {
	log.Println("Error hashing password:", err)
}
// Assume we've stored the hash securely from previous step
err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
if err != nil {
	log.Println("Hashed password does not match original password:", err)
} else {
	log.Println("Passwords match")
}
```

## Testing
The provided code snippet includes a basic test case for verifying the `Hash` function's functionality. It checks if the hash generated from the plaintext password matches the original password when comparing hashes using the `bcrypt.CompareHashAndPassword` function.

```go
package main

import (
	"testing"

	"yourproject/hasher" // Replace 'yourproject' with your project name
)

func TestHash(t *testing.T) {
	password := "password123"
	hash, err := hasher.Hash(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Hashed password does not match original password")
	}
}
```

