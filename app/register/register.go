package register

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/kilianp07/lifegame/database"
	"github.com/kilianp07/lifegame/model"
)

// Start is a function that starts the registration process by prompting the user for information and creating a new user in the database.
// It takes a controller parameter of type *database.Controller, which is used to interact with the database.
// It returns an error if there was an issue with prompting the user or creating the user in the database.
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

// completer is a function that takes a prompt.Document as input and returns a slice of prompt.Suggest.
// It is responsible for generating autocompletion suggestions based on the current input.
// The function filters the suggestions based on the prefix of the word before the cursor.
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// promptForUser is a function that prompts the user to enter their username and password.
// It returns a pointer to a User struct and an error if any.
func promptForUser() (*model.User, error) {
	var (
		User = &model.User{}
	)

	fmt.Println("Type your Username.")
	User.Username = prompt.Input("> ", completer)

	fmt.Println("Type your Password.")
	User.Password = prompt.Input("> ", completer)

	return User, User.Hash()
}
