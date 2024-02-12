 # Gameoflife - A Command-line Application for Conway's Game of Life

## Overview
`gameoflife` is a command-line application built in Go that simulates Conway's Game of Life, an intriguing cellular automaton devised by mathematician John Conway. The program provides users with the flexibility to define initial parameters for the game either through command line options or via a JSON configuration file.

## Features
- Customizable grid size and initial states.
- Ability to load configurations from a JSON file using the `--conf` flag.
- Real-time simulation of the game in the command line interface.

## Usage
```sh
gameoflife [options] [configuration_file]
```
### Options
| Short | Long    | Description                           |
|-------|---------|---------------------------------------|
| `-c`, `--conf`  | `Path to the JSON configuration file` | A JSON file containing the game's initial configuration. |

## Configuration File
A sample configuration file structure is as follows:
```json
{
	"grid_size": 5,
	"initial_state": [
		[1, 0],
		[0, 1],
		...
	]
}
```
The `grid_size` key specifies the dimensions of the grid, while the `initial_state` key is an array containing nested arrays representing each row in the grid. A value of `1` in this context represents a live cell, and `0` denotes a dead cell.

## Code Structure
```go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kilianp07/lifegame/app"
	"github.com/kilianp07/lifegame/config"
	"github.com/spf13/cobra"
)
```
The code is organized under the `cmd` package, which includes the main application logic for parsing command line arguments and configuration files, along with the `app` and `config` packages.

## Package Components
### `cmd`
This package contains the root command and its initialization, flag handling, and execution.
```go
import "github.com/spf13/cobra"

// ...

func init() {
	// Persistent flags
	rootCmd.PersistentFlags().StringVarP(&confPath, "conf", "c", "", "Path to the JSON configuration file")
}
```
### `app`
The application-specific logic for running the simulation is implemented in this package.
```go
type App struct {
	grid *Grid
}

func NewApp(config *config.App) *App {
	// Initialize grid
	grid := Grid{Size: config.GridSize, Matrix: make([][]int, config.GridSize)}

	// Set up initial state
	for i := range grid.Matrix {
		copy(grid.Matrix[i], config.InitialState[i])
	}

	return &App{Grid: &grid}
}

func (a *App) Start() {
	// Simulation logic goes here
}
```
### `config`
The configuration package is responsible for reading and deserializing JSON files.
```go
type App struct {
	GridSize int    `json:"grid_size"`
	InitialState [][]int `json:"initial_state"`
}

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
```

