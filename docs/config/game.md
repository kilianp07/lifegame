 # config Package

The `config` package is designed to store and manage configuration data for various applications. This simple Go structure provides a flexible way to define and decode JSON-encoded configuration files.

## Game Struct

```go
type Game struct {
	Lenght         int    `json:"lenght"`
	Height          int    `json:"height"`
	Generations     int    `json:"generations"`
	CharTrue        string `json:"char_true"`
	CharFalse       string `json:"char_false"`
	Sleep           string `json:"sleep"`
}
```

The `Game` struct is the primary data type in this package, which encapsulates essential configuration settings for a game. Below is an explanation of each field:

### Lenght

Type: int
JSON Tag: `lenght`

Represents the length of the game board or grid.

### Height

Type: int
JSON Tag: `height`

Defines the height of the game board or grid.

### Generations

Type: int
JSON Tag: `generations`

Specifies the number of generations for evolution simulations or similar processes in the game.

### CharTrue

Type: string
JSON Tag: `char_true`

Represents a character that symbolizes "true" or "alive" entities within the game's context.

### CharFalse

Type: string
JSON Tag: `char_false`

Symbolizes "false" or "dead" entities within the game's context.

### Sleep

Type: string
JSON Tag: `sleep`

Defines the duration of time that an entity, process, or game state will remain inactive or idle before moving on to the next step.

This configuration data is essential for applications like evolution simulations, cellular automata, or other games where settings need to be tailored to individual use cases. With this flexible and extensible configuration data structure, you'll have a powerful toolset at your disposal to build customizable, dynamic, and engaging applications in Go.

