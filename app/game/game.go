package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kilianp07/lifegame/config"
)

// Cell represents a single cell in the grid.
type Cell bool

// Game represents the state of the Game of Life.
type Game struct {
	Grid [][]Cell
	Rows int
	Cols int
}

// Constants for the cell state.
const (
	Alive Cell = true
	Dead  Cell = false
)

// NewGame creates a new game with the specified number of rows and columns.
func NewGame(rows, cols int) *Game {
	grid := make([][]Cell, rows)
	for i := range grid {
		grid[i] = make([]Cell, cols)
	}
	return &Game{Grid: grid, Rows: rows, Cols: cols}
}

// Initialize randomly initializes the grid with alive and dead cells.
func (g *Game) Initialize() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			g.Grid[i][j] = Cell(rand.Intn(2) == 1)
		}
	}
}

// NextState computes the next state of the game.
func (g *Game) NextState() {
	newGrid := make([][]Cell, g.Rows)
	for i := range newGrid {
		newGrid[i] = make([]Cell, g.Cols)
	}

	// Calculate the next state based on the current grid state
	for x := 0; x < g.Rows; x++ {
		for y := 0; y < g.Cols; y++ {
			aliveNeighbors := g.countAliveNeighbors(x, y)
			switch {
			case g.Grid[x][y] == Dead && aliveNeighbors == 3:
				newGrid[x][y] = Alive
			case g.Grid[x][y] == Alive && (aliveNeighbors == 2 || aliveNeighbors == 3):
				newGrid[x][y] = Alive
			default:
				newGrid[x][y] = Dead
			}
		}
	}

	// Update the grid with the new state after all calculations are complete
	g.Grid = newGrid
}

// countAliveNeighbors counts the alive neighbors of a cell.
func (g *Game) countAliveNeighbors(x, y int) int {
	neighbors := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && ny >= 0 && nx < g.Rows && ny < g.Cols && g.Grid[nx][ny] == Alive {
				neighbors++
			}
		}
	}
	return neighbors
}

// Display prints the grid to the console.
func (g *Game) Display(chartTrue, chartFalse string) {
	for _, row := range g.Grid {
		for _, cell := range row {
			if cell == Alive {
				fmt.Print(chartTrue, " ")
			} else {
				fmt.Print(chartFalse, " ")
			}
		}
		fmt.Println()
	}
}

func Run(conf config.Game) {
	game := NewGame(conf.Height, conf.Lenght)
	fmt.Println(conf)
	game.Initialize()

	sleep, err := time.ParseDuration(conf.Sleep)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < conf.Generations; i++ {

		game.Display(conf.CharTrue, conf.CharFalse)
		fmt.Println("-----")
		game.NextState()
		time.Sleep(sleep)
	}

}
