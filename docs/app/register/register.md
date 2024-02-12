 # register Package

The `register` package is a Go package designed for user registration functionality. It imports required dependencies and defines functions to handle user input and database interactions.

## Dependencies

```go
import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/kilianp07/lifegame/database"
	"github.com/kilianp07/lifegame/model"
)
```

## Functions

### `Start(controller *database.Controller) error`

The `Start` function initializes the registration process by requesting user input and creating a new entry in the database using the provided controller. It takes as an argument a pointer to a `database.Controller`, which is used to interact with the underlying database. The function returns an error if there were any issues with prompting the user or creating the user record in the database.

```go
func Start(controller *database.Controller) error {
	var (
		user *model.User
		err  error
	)

	user, err = promptForUser()
	if err != nil {
		return err
	}

	return controller.CreateUser(user)
}
```

### `completer(d prompt.Document) []prompt.Suggest`

The `completer` function generates autocompletion suggestions based on the current input when using a Go-Prompt terminal. It filters these suggestions based on the prefix of the word before the cursor.

```go
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
```

### `promptForUser() (*model.User, error)`

The `promptForUser()` function requests the user to enter their username and password through a terminal interface. It returns a pointer to a new instance of the `model.User` struct and any potential errors that might occur during the input process (e.g., if the user fails to provide valid input).

```go
func promptForUser() (*model.User, error) {
	var (
		User = &model.User{}
	)

	fmt.Println("Type your Username.")
	User.Username, _ = prompt.Input("> ", completer)

	fmt.Println("Type your Password.")
	User.Password, err = prompt.Input("> ", nil) // No need for completer as we don't want suggestions here

	return User, User.Hash()
}
```

