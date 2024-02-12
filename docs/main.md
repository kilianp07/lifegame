 # Golang Project Documentation

## Overview
This documentation describes a simple GoLang project with the name `lifegame`. The project is structured as follows:

```go
package main
import (
	"github.com/kilianp07/lifegame/cmd"
)

func main() {
	cmd.Execute()
}
```

## Project Structure
The project consists of the following components:
- **main.go**: Entry point of the application. It imports and initializes the `cmd` package from `lifegame/cmd`. The `main` function calls the `Execute` method of the `cmd.App` instance, which starts the application.
- **github.com/kilianp07/lifegame/cmd**: This is a sub-package that contains the application's main entry point and commands for interacting with the game. The `Execute` function parses command line arguments, sets up configuration options, and calls the appropriate command handler based on the provided command.

## Getting Started
To get started, you can clone this repository and install the project using GoMod:
```sh
$ git clone https://github.com/username/repository.git
$ cd lifegame
$ go mod init github.com/username/lifegame
```

After setting up the project, you can build and run the application using the following commands:
```sh
$ go build -o lifegame
$ ./lifegame [options]
```

## Usage
For a detailed list of available options and their usage, you can use the `--help` flag when running the application:
```sh
$ ./lifegame --help
```

This command will print out the complete usage information, including all available commands and their flags.

## Dependencies
The project has a dependency on the `github.com/kilianp07/lifegame/cmd` package which is used to implement the application's logic. Currently, there are no other dependencies declared in the project's `go.mod` file.

## License
This project is licensed under the MIT license - see the [LICENSE](LICENSE) file for details.

