package app

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/kilianp07/lifegame/app/register"
	"github.com/kilianp07/lifegame/config"
	"github.com/kilianp07/lifegame/database"
)

const (
	REGISTER = "register"
	PLAY     = "play"
	RULES    = "rules"
	EXIT     = "exit"
)

type App struct {
	conf       *config.App
	controller *database.Controller
}

func NewApp(conf *config.App) *App {
	return &App{
		conf: conf,
	}
}

// Start is a function that starts the application.
// It prompts the user to select a table and calls the corresponding function.
// The function continues to prompt the user until the choice is "exit".
// If an error occurs while determining the function for the selected table,
// the error is printed and the loop is exited.
func (a *App) Start() {
	var (
		choice string
		err    error
	)

	if err = a.configure(); err != nil {
		fmt.Println(err)
		return
	}

	for choice != "exit" {
		fmt.Println("Please select table.")
		choice = prompt.Input("> ", a.completer)
		if err = a.determineFunc(choice); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func (a *App) configure() error {
	var (
		err error
	)

	a.controller, err = database.NewController(a.conf.DBConf)
	if err != nil {
		return err
	}

	return a.controller.Init()
}

// completer is a function that provides autocompletion suggestions for the prompt.
// It takes a prompt.Document as input and returns a slice of prompt.Suggest.
// The suggestions include options for creating a new user account, playing the game,
// reading the rules, and exiting the game.
// The suggestions are filtered based on the word before the cursor in the input document.
func (a *App) completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: REGISTER, Description: "Create a new user account"},
		{Text: PLAY, Description: "Play the game"},
		{Text: RULES, Description: "Read rules"},
		{Text: EXIT, Description: "Exit the game"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// determineFunc is a function that takes a string as input and performs a specific action based on the value of the string.
// It returns an error if an error occurs during the execution of the action.
func (a *App) determineFunc(s string) error {
	switch s {
	case REGISTER:
		return register.Start(a.controller)
	case PLAY:
		fmt.Println("play")
	case RULES:
		fmt.Println("rules")
	case EXIT:
		break
	default:
		fmt.Println("unknown command")
	}
	return nil
}
