package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var input []byte
	var err error

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		input, err = os.ReadFile("input.txt")
		if err != nil {
			panic(err)
		}
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}

func part1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])
	accessible := 0

	// directions for 8 adjacent cells (diagonals + orthogonal)
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// check each cell in the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if lines[r][c] != '@' {
				continue
			}

			// count adjacent paper rolls
			adjacentCount := 0
			for _, dir := range directions {
				nr := r + dir[0]
				nc := c + dir[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if lines[nr][nc] == '@' {
						adjacentCount++
					}
				}
			}

			// forklift can access if fewer than 4 adjacent rolls
			if adjacentCount < 4 {
				accessible++
			}
		}
	}

	return accessible
}

func part2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])

	// create a mutable grid
	grid := make([][]byte, rows)
	for i := range lines {
		grid[i] = []byte(lines[i])
	}

	totalRemoved := 0

	// directions for 8 adjacent cells
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// repeatedly remove accessible rolls until none remain
	// removing rolls can make other rolls become accessible
	for {
		// find all currently accessible rolls
		accessible := [][]int{}

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] != '@' {
					continue
				}

				// count adjacent paper rolls
				adjacentCount := 0
				for _, dir := range directions {
					nr := r + dir[0]
					nc := c + dir[1]

					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						if grid[nr][nc] == '@' {
							adjacentCount++
						}
					}
				}

				// accessible if fewer than 4 adjacent rolls
				if adjacentCount < 4 {
					accessible = append(accessible, []int{r, c})
				}
			}
		}

		// if no more accessible rolls, stop the loop
		if len(accessible) == 0 {
			break
		}

		// remove all accessible rolls in this iteration
		for _, pos := range accessible {
			grid[pos[0]][pos[1]] = '.'
		}

		totalRemoved += len(accessible)
	}

	return totalRemoved
}

