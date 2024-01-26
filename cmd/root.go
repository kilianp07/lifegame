package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kilianp07/lifegame/app"
	"github.com/kilianp07/lifegame/config"
	"github.com/spf13/cobra"
)

var (
	// confPath is the variable where the configuration file path will be stored
	confPath string

	// rootCmd represents the root command of the application
	rootCmd = &cobra.Command{
		Use:   "gameoflife",
		Short: "Simulate Conway's Game of Life in your command line",
		Long: `Gameoflife is a command-line application that simulates Conway's Game of Life, an intriguing cellular automaton devised by mathematician John Conway. This program allows users to define initial parameters for the game, such as grid size and the initial state of cells, either directly through command line options or via a JSON configuration file.
		
		Features include:
		- Configurable grid size and initial states.
		- Option to load configurations from a JSON file using the --conf flag.
		- Real-time simulation of the game in the command line interface.
		
		Use --help to see the available commands and options.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Your main command logic goes here
			if confPath != "" {
				config, err := loadConfig(confPath)
				if err != nil {
					fmt.Println("Error loading configuration file:", err)
					os.Exit(1)
				}
				// Use the loaded configuration here
				fmt.Println("Configuration loaded")
				app.NewApp(config).Start()
			} else {
				fmt.Println("No configuration file specified, parameter must be filled to start app.")
				return cmd.Help()
			}
			return nil
		},
	}
)

// loadConfig reads and deserializes the JSON configuration file.
func loadConfig(confPath string) (*config.App, error) {
	file, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := config.App{}
	err = decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// init initializes the flags and configuration for the root command.
func init() {
	rootCmd.PersistentFlags().StringVarP(&confPath, "conf", "c", "", "Path to the JSON configuration file")
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
