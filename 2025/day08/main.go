package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

type point struct {
	x, y, z int
}

type pair struct {
	i, j   int
	distSq int64
}

func parsePoints(lines []string) []point {
	points := make([]point, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, point{x, y, z})
	}
	return points
}

func distSq(a, b point) int64 {
	dx := int64(a.x - b.x)
	dy := int64(a.y - b.y)
	dz := int64(a.z - b.z)
	return dx*dx + dy*dy + dz*dz
}

func getSortedPairs(points []point) []pair {
	n := len(points)
	pairs := make([]pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, pair{i, j, distSq(points[i], points[j])})
		}
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].distSq < pairs[b].distSq
	})
	return pairs
}

// union-find data structure
type unionFind struct {
	parent []int
	rank   []int
	size   []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, rank, size}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	uf.size[px] += uf.size[py]
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	return true
}

func part1(lines []string) int {
	points := parsePoints(lines)
	n := len(points)
	pairs := getSortedPairs(points)

	uf := newUnionFind(n)

	// connect the 1000 closest pairs
	for i := 0; i < 1000 && i < len(pairs); i++ {
		uf.union(pairs[i].i, pairs[i].j)
	}

	// collect circuit sizes
	sizes := make(map[int]int)
	for i := 0; i < n; i++ {
		root := uf.find(i)
		sizes[root] = uf.size[root]
	}

	// get unique sizes and sort descending
	sizeList := make([]int, 0, len(sizes))
	for _, s := range sizes {
		sizeList = append(sizeList, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizeList)))

	// multiply top 3
	result := 1
	for i := 0; i < 3 && i < len(sizeList); i++ {
		result *= sizeList[i]
	}

	return result
}

func part2(lines []string) int {
	points := parsePoints(lines)
	n := len(points)
	pairs := getSortedPairs(points)

	uf := newUnionFind(n)
	numComponents := n

	// keep connecting until all in one circuit
	var lastPair pair
	for _, p := range pairs {
		if uf.union(p.i, p.j) {
			numComponents--
			lastPair = p
			if numComponents == 1 {
				break
			}
		}
	}

	// return product of x coordinates
	return points[lastPair.i].x * points[lastPair.j].x
}
