package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println("part 1:", part1(lines))
	fmt.Println("part 2:", part2(lines))
}

func readInput() ([]byte, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile("input.txt")
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}

