package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

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

	content := strings.TrimSpace(string(input))
	parts := strings.Split(content, "\n\n")

	if len(parts) != 2 {
		panic("invalid input format: expected ranges and ingredients separated by blank line")
	}

	ranges := parseRanges(parts[0])
	ingredients := parseIngredients(parts[1])

	fmt.Println("part 1:", part1(ranges, ingredients))
	fmt.Println("part 2:", part2(ranges))
}

func parseRanges(section string) []Range {
	lines := strings.Split(section, "\n")
	ranges := make([]Range, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic("invalid range format: " + line)
		}

		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic("invalid start value: " + parts[0])
		}

		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic("invalid end value: " + parts[1])
		}

		ranges = append(ranges, Range{start: start, end: end})
	}

	return ranges
}

func parseIngredients(section string) []int64 {
	lines := strings.Split(section, "\n")
	ingredients := make([]int64, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic("invalid ingredient id: " + line)
		}

		ingredients = append(ingredients, id)
	}

	return ingredients
}

func isFresh(id int64, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func part1(ranges []Range, ingredients []int64) int {
	freshCount := 0

	for _, id := range ingredients {
		if isFresh(id, ranges) {
			freshCount++
		}
	}

	return freshCount
}

func part2(ranges []Range) int64 {
	if len(ranges) == 0 {
		return 0
	}

	// sort ranges by start value
	sorted := make([]Range, len(ranges))
	copy(sorted, ranges)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].start < sorted[j].start
	})

	// merge overlapping ranges
	merged := []Range{sorted[0]}

	for i := 1; i < len(sorted); i++ {
		current := sorted[i]
		last := &merged[len(merged)-1]

		// ranges overlap or are adjacent if current.start <= last.end + 1
		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}

	// count total unique IDs
	var total int64 = 0
	for _, r := range merged {
		total += r.end - r.start + 1
	}

	return total
}

