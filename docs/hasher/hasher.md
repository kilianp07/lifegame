 # package hasher

This package provides functions for password hashing and verification using the `bcrypt` library from the `golang.org/x/crypto` package. The primary goal is to ensure secure handling of user passwords in a Go application.

## Importing the Package

```go
import "golang.org/x/crypto/bcrypt"

// or
import hasher "path/to/your/hasher_package"
```

## Functions

### 1. Hash Function

`Hash` is a function that takes a plain text password string as an argument and returns its hashed representation. It utilizes the `bcrypt.GenerateFromPassword()` function to generate a secure hash with a cost factor of 14.

```go
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
```

#### Input
- `password`: A plain text password string

#### Output
- `hashedPassword`: A hashed representation of the given password as a string
- `error`: An error value if an issue occurs during the hashing process; otherwise, it will be `nil`

### 2. CheckPasswordHash Function

`CheckPasswordHash` is a function that takes two string parameters - a plain text password and its corresponding hash. It compares the provided password with the stored hash and returns a boolean value indicating if they match. The comparison is performed using the `bcrypt.CompareHashAndPassword()` function.

```go
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
```

#### Input
- `password`: A plain text password string to be verified
- `hash`: The stored hashed representation of the corresponding password as a string

#### Output
- `true` if the plain text password matches its hash; otherwise, `false`

## Example Usage

```go
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// ... Use your import statement here instead of this one

func main() {
	plainTextPassword := "password123!"
	hashedPassword, _ := Hash(plainTextPassword)
	fmt.Println("Hashed password:", hashedPassword)

	if CheckPasswordHash(plainTextPassword, hashedPassword) {
		fmt.Println("Password matches the hash.")
	} else {
		fmt.Println("Password does not match the hash.")
	}
}
```

