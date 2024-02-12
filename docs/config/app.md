 # config Package

The `config` package is designed to handle the configuration of our application in Go. This documentation will explain the structure and usage of the package, which includes an `App` type and a nested `DBConf` and `Game` types.

## App Type

```go
type App struct {
	DBConf   DBConf   `json:"db"`
	Game     Game     `json:"game"`
}
```

The main data structure of the package is an `App` type, which holds two nested structures: `DBConf` and `Game`. These nested types will be explained in detail below.

### DBConf Type

```go
type DBConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
```

The `DBConf` type represents the database configuration for our application. It contains the following fields:

- `Host`: The host address of the database.
- `Port`: The port number for connecting to the database.
- `User`: The username for accessing the database.
- `Password`: The password for accessing the database.
- `Name`: The name of the database that our application will connect to.

### Game Type

```go
type Game struct {
	Mode   string `json:"mode"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
```

The `Game` type represents the configuration for our application's game component. It contains the following fields:

- `Mode`: The current game mode (e.g., "singleplayer" or "multiplayer").
- `Width`: The width of the game window in pixels.
- `Height`: The height of the game window in pixels.

## Usage

To use the `config` package, first import it into your Go project:

```go
import (
	"github.com/yourusername/config"
)
```

Next, create a new instance of the `App` type and parse your configuration file:

```go
func main() {
	app := config.App{}
	err := json.Unmarshal([]byte(configFile), &app)
	if err != nil {
		log.Fatal("Error unmarshalling config:", err)
	}
	// Continue with the parsed configuration...
}
```

Replace `configFile` with the contents of your JSON-formatted configuration file. Once the configuration has been successfully parsed, you can access it through the various fields of the `App`, `DBConf`, and `Game` structures.

