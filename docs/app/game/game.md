 # game Package Documentation

The `game` package is a simple implementation of the Game of Life, a cellular automaton devised by the British mathematician John Horton Conway in 1970. This simulation explores how life evolves according to its inherent rules.

## Table of Contents

- [Types](#types)
  * [`Cell`](#cell)
  * [`Game`](#game)
- [Constants](#constants)
- [Functions](#functions)
  * [`NewGame`](#newgame)
  * [`Initialize`](#initialize)
  * [`NextState`](#nextstate)
  * [`countAliveNeighbors`](#countalivenebors)
  * [`Display`](#display)
  * [`Run`](#run)

## Types <a name="types"></a>

### Cell <a name="cell"></a>

A `Cell` is a boolean type representing the state of a cell in the grid.

```go
type Cell bool
```

#### Constants

Two constants define possible states for a `Cell`.

```go
const (
	Alive  Cell = true
	Dead   Cell = false
)
```

### Game <a name="game"></a>

A `Game` is a struct that holds the state of the Game of Life, consisting of a grid with cells and its dimensions.

```go
type Game struct {
	Grid [][]Cell
	Rows int
	Cols int
}
```

## Constants <a name="constants"></a>

### `Alive` and `Dead`

Two constants are defined for the cell state.

```go
const (
	Alive Cell = true
	Dead  Cell = false
)
```

## Functions <a name="functions"></a>

### `NewGame` <a name="newgame"></a>

The `NewGame` function creates a new game with the specified number of rows and columns.

```go
func NewGame(rows, cols int) *Game {
	// ... initialization logic here ...
}
```

### `Initialize` <a name="initialize"></a>

The `Initialize` method randomly initializes the grid with alive and dead cells.

```go
func (g *Game) Initialize() {
	// ... randomization logic here ...
}
```

### `NextState` <a name="nextstate"></a>

The `NextState` method computes the next state of the game based on the current grid state.

```go
func (g *Game) NextState() {
	// ... next state calculation logic here ...
}
```

### `countAliveNeighbors` <a name="countalivenebors"></a>

The `countAliveNeighbors` method counts the number of alive neighbors around a cell.

```go
func (g *Game) countAliveNeighbors(x, y int) int {
	// ... counting logic here ...
}
```

### `Display` <a name="display"></a>

The `Display` method prints the grid to the console.

```go
func (g *Game) Display(chartTrue, chartFalse string) {
	// ... display logic here ...
}
```

### `Run` <a name="run"></a>

The `Run` function initializes a new game and simulates it for the specified number of generations.

```go
func Run(conf config.Game) {
	// ... initialization, simulation, and display logic here ...
}
```

