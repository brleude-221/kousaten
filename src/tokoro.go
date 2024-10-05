package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Crossroad represents each intersection with its coordinates and traffic light timings
type Crossroad struct {
	X, Y          int  // Coordinates of the crossroad
	RedDuration   int  // Duration of red light
	GreenDuration int  // Duration of green light
	FirstDuration int  // Initial random duration (first cycle)
	FirstIsRed    bool // Whether the first color is red (otherwise green)
}

// Board represents the game board with size and crossroad details
type Board struct {
	Size  int
	Grid  [][]Crossroad
	Start Crossroad
	End   Crossroad
}

// NewCrossroad creates a crossroad with random light timings at given coordinates
func NewCrossroad(x, y int) Crossroad {
	red := rand.Intn(10) + 5    // Random red light duration between 5 and 15 seconds
	green := rand.Intn(10) + 5  // Random green light duration between 5 and 15 seconds
	firstIsRed := rand.Intn(2) == 0             // Randomly decide if the first cycle starts with red or green

  var firstDuration int
	if firstIsRed {
	  // If the first light is red, ensure the first duration is not longer than the red duration
		firstDuration = rand.Intn(red) + 1
  	} else {
	  // If the first light is green, ensure the first duration is not longer than the green duration
	  firstDuration = rand.Intn(green) + 1
  }
	return Crossroad{
		X:            x,
		Y:            y,
		RedDuration:  red,
		GreenDuration: green,
		FirstDuration: firstDuration,
		FirstIsRed:   firstIsRed,
	}
}

// randomCrossroad selects a random crossroad on the board
func randomCrossroad(size int) Crossroad {
	x := rand.Intn(size)
	y := rand.Intn(size)
	return NewCrossroad(x, y)
}

// Not Sure if I like this function...
// randomEndPoint selects an end point that is strictly to the right and up of the start point
func randomEndPoint(size int, start Crossroad) Crossroad {
	var end Crossroad
	for {
		end = randomCrossroad(size)
		// Ensure the end point is both strictly to the right (X > start.X) and up (Y > start.Y)
		if end.X > start.X && end.Y > start.Y {
			break
		}
	}
	return end
}

// NewBoard creates a new board of the given size and randomly selects start and end points
func NewBoard(size int) *Board {
	rand.Seed(time.Now().UnixNano())

	grid := make([][]Crossroad, size)
	for i := range grid {
		grid[i] = make([]Crossroad, size)
		for j := range grid[i] {
			grid[i][j] = NewCrossroad(i, j)
		}
	}

	start := randomCrossroad(size)
	end := randomEndPoint(size, start)

	return &Board{
		Size:  size,
		Grid:  grid,
		Start: start,
		End:   end,
	}
}

// DisplayBoard prints the board and crossroad details
func (b *Board) DisplayBoard() {
	fmt.Printf("Board size: %dx%d\n", b.Size, b.Size)
	fmt.Printf("Start Point: (%d, %d)\n", b.Start.X, b.Start.Y)
	fmt.Printf("End Point: (%d, %d)\n", b.End.X, b.End.Y)
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			c := b.Grid[i][j]
			fmt.Printf("(%d,%d): Red=%ds, Green=%ds, First=%ds, Starts with %s\n",
				c.X, c.Y, c.RedDuration, c.GreenDuration, c.FirstDuration, 
				map[bool]string{true: "Red", false: "Green"}[c.FirstIsRed])
		}
	}
}

func main() {
	// Create a board of size 10x10 (can be changed as needed)
	board := NewBoard(10)

	// Display the generated board and crossroad details
	board.DisplayBoard()
}

