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
	if len(lines) < 2 {
		return 0
	}
	
	numRows := len(lines) - 1
	opRow := len(lines) - 1
	
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	
	// pad all lines to same length for column-based parsing
	padded := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		padded[i] = lines[i]
		for len(padded[i]) < maxLen {
			padded[i] = padded[i] + " "
		}
	}
	
	problems := []struct {
		numbers []int
		op      rune
	}{}
	
	// parse problems by scanning columns left to right
	col := 0
	for col < maxLen {
		isEmptyCol := true
		for row := 0; row < len(lines); row++ {
			if padded[row][col] != ' ' {
				isEmptyCol = false
				break
			}
		}
		
		if isEmptyCol {
			col++
			continue
		}
		
		startCol := col
		for col < maxLen {
			hasNonSpace := false
			for row := 0; row < len(lines); row++ {
				if padded[row][col] != ' ' {
					hasNonSpace = true
					break
				}
			}
			if hasNonSpace {
				col++
			} else {
				break
			}
		}
		
		endCol := col
		
		numbers := []int{}
		var operation rune
		
		for row := 0; row < numRows; row++ {
			numStr := ""
			for c := startCol; c < endCol; c++ {
				if padded[row][c] != ' ' {
					numStr = numStr + string(padded[row][c])
				}
			}
			if numStr != "" {
				num := 0
				for _, ch := range numStr {
					num = num*10 + int(ch-'0')
				}
				numbers = append(numbers, num)
			}
		}
		
		for c := startCol; c < endCol; c++ {
			if padded[opRow][c] != ' ' {
				operation = rune(padded[opRow][c])
				break
			}
		}
		
		if len(numbers) > 0 && (operation == '*' || operation == '+') {
			problems = append(problems, struct {
				numbers []int
				op      rune
			}{numbers, operation})
		}
	}
	
	grandTotal := 0
	for _, prob := range problems {
		result := prob.numbers[0]
		for i := 1; i < len(prob.numbers); i++ {
			if prob.op == '*' {
				result = result * prob.numbers[i]
			} else {
				result = result + prob.numbers[i]
			}
		}
		grandTotal = grandTotal + result
	}
	
	return grandTotal
}

func part2(lines []string) int {
	if len(lines) < 2 {
		return 0
	}
	
	numRows := len(lines) - 1
	opRow := len(lines) - 1
	
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	
	// pad all lines to same length for column-based parsing
	padded := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		padded[i] = lines[i]
		for len(padded[i]) < maxLen {
			padded[i] = padded[i] + " "
		}
	}
	
	problems := []struct {
		numbers []int
		op      rune
	}{}
	
	// parse problems by scanning columns, but read each column in reverse order
	col := 0
	for col < maxLen {
		isEmptyCol := true
		for row := 0; row < len(lines); row++ {
			if padded[row][col] != ' ' {
				isEmptyCol = false
				break
			}
		}
		
		if isEmptyCol {
			col++
			continue
		}
		
		startCol := col
		for col < maxLen {
			hasNonSpace := false
			for row := 0; row < len(lines); row++ {
				if padded[row][col] != ' ' {
					hasNonSpace = true
					break
				}
			}
			if hasNonSpace {
				col++
			} else {
				break
			}
		}
		
		endCol := col
		
		numbers := []int{}
		var operation rune
		
		for c := endCol - 1; c >= startCol; c-- {
			numStr := ""
			for row := 0; row < numRows; row++ {
				if padded[row][c] != ' ' {
					numStr = numStr + string(padded[row][c])
				}
			}
			
			if numStr != "" {
				num := 0
				for _, ch := range numStr {
					num = num*10 + int(ch-'0')
				}
				numbers = append(numbers, num)
			}
			
			if padded[opRow][c] != ' ' && operation == 0 {
				operation = rune(padded[opRow][c])
			}
		}
		
		if len(numbers) > 0 && (operation == '*' || operation == '+') {
			problems = append(problems, struct {
				numbers []int
				op      rune
			}{numbers, operation})
		}
	}
	
	grandTotal := 0
	for _, prob := range problems {
		result := prob.numbers[0]
		for i := 1; i < len(prob.numbers); i++ {
			if prob.op == '*' {
				result = result * prob.numbers[i]
			} else {
				result = result + prob.numbers[i]
			}
		}
		grandTotal = grandTotal + result
	}
	
	return grandTotal
}

