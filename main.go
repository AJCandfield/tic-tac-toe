package main

import (
	"container/ring"
	"fmt"
)

func main() {

	// Create 3x3 grid
	grid := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	// Create a ring buffer to host the player's symbols
	r := ring.New(2)

	r.Value = "x"
	r = r.Next()
	r.Value = "o"

	for {
		var player string

		// Get the current player's symbol
		player = r.Value.(string)

		var xCord int
		var yCord int

		println("Current playa:", player)

		fmt.Print("Enter position (x): ")
		fmt.Scan(&xCord)

		fmt.Print("Enter position (y): ")
		fmt.Scan(&yCord)

		if grid[yCord-1][xCord-1] != " " {
			println("Illegal move!")
			continue
		}

		grid[yCord-1][xCord-1] = player

		// For each row
		for y := 0; y < len(grid); y++ {
			// For each column
			for x := 0; x < len(grid[y]); x++ {
				fmt.Print("[", grid[y][x], "]")
			}
			println()
		}
		r = r.Next()
	}
}
