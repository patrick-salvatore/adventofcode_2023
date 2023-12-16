package main

import (
	"fmt"
	"os"
	"strings"
)

func get_starting_coor(grid [][]string) (int, int) {
	var start_r, start_c int

	for r, row := range grid {
		for c, col := range row {
			if col == "S" {
				start_r = r
				start_c = c
				return start_r, start_c
			}
		}
	}

	return -1, -1
}

type Coor struct {
	r, c int
}

func is_one_of(ch string, s string) bool {

	for _, c := range s {
		if ch == string(c) {
			return true
		}
	}
	return false
}

func part1(lines []string) float64 {
	var grid = [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	start_r, start_c := get_starting_coor(grid)

	visited := make(map[string]bool)
	visited[fmt.Sprintf("%v,%v", start_r, start_c)] = true

	var q = []Coor{{r: start_r, c: start_c}}

	for len(q) > 0 {
		fmt.Println(q)
		coor := q[0]
		q = q[1:]
		ch := grid[coor.r][coor.c]

		// Can we go up?
		if is_one_of(ch, "S|LJ") && coor.r > 0 && is_one_of(grid[coor.r-1][coor.c], "|7F") && !visited[fmt.Sprintf("%v,%v", coor.r-1, coor.c)] {
			q = append(q, Coor{r: coor.r - 1, c: coor.c})
			visited[fmt.Sprintf("%v,%v", coor.r-1, coor.c)] = true
		}

		// Can we go down?
		if is_one_of(ch, "S|7F") && coor.r < len(grid)-1 && is_one_of(grid[coor.r+1][coor.c], "|LJ") && !visited[fmt.Sprintf("%v,%v", coor.r+1, coor.c)] {
			q = append(q, Coor{r: coor.r + 1, c: coor.c})
			visited[fmt.Sprintf("%v,%v", coor.r+1, coor.c)] = true
		}

		// Can we go to the left?
		if is_one_of(ch, "S-J7") && coor.c > 0 && is_one_of(grid[coor.r][coor.c-1], "-LF") && !visited[fmt.Sprintf("%v,%v", coor.r, coor.c-1)] {
			q = append(q, Coor{r: coor.r, c: coor.c - 1})
			visited[fmt.Sprintf("%v,%v", coor.r, coor.c-1)] = true
		}

		// Can we go to the right?
		if is_one_of(ch, "S-LF") && coor.c < len(grid[coor.r])-1 && is_one_of(grid[coor.r][coor.c+1], "-J7") && !visited[fmt.Sprintf("%v,%v", coor.r, coor.c+1)] {
			q = append(q, Coor{r: coor.r, c: coor.c + 1})
			visited[fmt.Sprintf("%v,%v", coor.r, coor.c+1)] = true
		}
	}
	return float64(len(visited) / 2)
}

func part2(lines []string) int {
	result := 0

	return result
}

func main() {
	data, err := os.ReadFile("day_10/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Result for part1: %v\n", part1(lines))
	fmt.Printf("Result for part2: %v\n", part2(lines))
}
