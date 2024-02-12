 # model Package Documentation

The `model` package is responsible for defining the data structures and handling data manipulation in the application. It imports necessary packages like `github.com/kilianp07/lifegame/hasher` and `gorm.io/gorm`.

## Structs

### User

The `User` struct represents a user account in our system. It extends the default GORM model with unique username and password fields:

```go
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}
```

- The `gorm.Model` type provides common fields for all models like `ID`, `CreatedAt`, `UpdatedAt`, and `DeletedAt`.
- The `Username` field is unique, which ensures no two users can have the same username.
- The `Password` field will be hashed using a secure hashing algorithm.

## Functions

### Hash

The `Hash()` function is defined on the User struct and is used to hash the user's password:

```go
func (u *User) Hash() error {
	hash, err := hasher.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}
```

This function performs the following actions:

1. Calls the `Hash()` function from the imported `hasher` package to hash the user's password.
2. If an error occurs during the hashing process, it is returned immediately.
3. In case of success, the user's `Password` field is updated with the new hash value and no error is returned.

