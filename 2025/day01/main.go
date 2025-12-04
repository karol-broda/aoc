package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
	position := 50
	count := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		// rotate the dial based on direction
		// L = left (counterclockwise, toward lower numbers)
		// R = right (clockwise, toward higher numbers)
		if direction == 'L' {
			position = (position - distance) % 100
			// handle wraparound for negative values (e.g., 5 - 10 = -5 -> 95)
			if position < 0 {
				position = position + 100
			}
		} else if direction == 'R' {
			position = (position + distance) % 100
		}

		// count every time dial points at 0
		if position == 0 {
			count++
		}
	}

	return count
}

func part2(lines []string) int {
	position := 50
	count := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		// count how many times the dial passes through 0 during rotation
		if direction == 'L' {
			// rotating left (counterclockwise)
			if position == 0 {
				// starting at 0, dial hits 0 again only after complete rotations
				count += distance / 100
			} else {
				// count complete rotations plus one more if the dial reaches 0 in a partial rotation
				count += distance / 100
				if distance%100 >= position {
					count++
				}
			}
			// update position
			position = (position - distance) % 100
			if position < 0 {
				position = position + 100
			}
		} else if direction == 'R' {
			// rotating right (clockwise)
			// count how many times dial crosses 0 boundary
			count += (position + distance) / 100
			// update position
			position = (position + distance) % 100
		}
	}

	return count
}
