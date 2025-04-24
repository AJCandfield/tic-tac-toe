package main

import (
	"container/ring"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("start of execution")
	// Create 3x3 grid.
	// Horizontal lines are contained in the vertical lines.
	//
	// Therefore, grid[y][x]
	//
	// Example: grid[0]    == the first row of squares
	// Example: grid[1][2] == the second row of squares, the third square
	grid := [][]string{
		// <--x-axis -->
		{" ", " ", " "}, // }
		{" ", " ", " "}, // } -> y-axis
		{" ", " ", " "}, // }
	}

	// Create a ring buffer to host the player's symbols
	r := ring.New(2)

	r.Value = "x"
	r = r.Next()
	r.Value = "o"

	for {
		// Get the current player's symbol
		player := r.Value.(string)
		log_turn := logger.With("player", player)
		log_turn.Debug("new turn started")

		// Draw grid with updated symbols
		// For each row
		for y := 0; y < len(grid); y++ {
			// For each column
			for x := 0; x < len(grid[y]); x++ {
				fmt.Print("[", grid[y][x], "]")
			}
			println()
		}
		log_turn.Debug("printed grid")

		var xCord int
		var yCord int

		println("Current player:", player)

		fmt.Print("Enter position (column): ")
		fmt.Scan(&xCord)

		fmt.Print("Enter position (row): ")
		fmt.Scan(&yCord)

		// Check if selected square is free
		if grid[yCord-1][xCord-1] != " " {
			println("Illegal move!")
			continue
		}

		// Assign player symbol to selected square
		grid[yCord-1][xCord-1] = player

		// === DETECT WIN ===
		//
		// there are 3 win conditions
		// - Any horizontal line (3)
		// - Any vertical line (3)
		// - Any diagonal line (2)
		//
		// We check the current symbol/player in each square
		// We break the check when the symbol/player doesn't match

		// Check horizontal win
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] != player {
					break
				}
				if x == len(grid[y])-1 {
					fmt.Println("Player", player, "won!")
				}
			}
		}

		// Check vertical win
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				if grid[x][y] != player {
					break
				}
				if x == len(grid[x])-1 {
					fmt.Println("Player", player, "won!")
				}
			}
		}

		// Check first diagonal win
		for z := 0; z < len(grid); z++ {
			if grid[z][z] != player {
				break
			}
			if z == len(grid)-1 {
				fmt.Println("Player", player, "won!")
			}
		}

		// Check second diagonal win
		for z := 0; z < len(grid); z++ {
			if grid[z][(len(grid)-1)-z] != player {
				break
			}
			if z == len(grid)-1 {
				fmt.Println("Player", player, "won!")
			}
		}

		// Call next turn
		r = r.Next()
	}
}
