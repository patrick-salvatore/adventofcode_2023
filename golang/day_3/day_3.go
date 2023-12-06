package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Coor struct {
	Row, Col int
}

func multiply(list []int) int {
	total := 1

	for _, num := range list {
		total *= num
	}

	return total
}

func sum_list(list []int) int {
	sum := 0

	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func is_symbol(s string) bool {
	match := regexp.MustCompile(`[^a-zA-Z0-9.]`).FindAllStringSubmatch(s, -1)

	return len(match) > 0
}

func get_neighbors_part_1(matrix []string, row, col int) []string {
	rowsLen := len(matrix) - 1
	colsLen := len(matrix[0]) - 1

	if rowsLen == 0 {
		return nil
	}

	positions := []Coor{
		{row - 1, col},     // top
		{row - 1, col + 1}, // top-right
		{row, col + 1},     // right
		{row + 1, col + 1}, // bottom-right
		{row + 1, col},     // bottom
		{row + 1, col - 1}, // bottom-left
		{row, col - 1},     // left
		{row - 1, col - 1}, // top-left
	}

	var validNeighbors []string
	for _, pos := range positions {
		if pos.Row >= 0 && pos.Row < rowsLen && pos.Col >= 0 && pos.Col < colsLen {
			validNeighbors = append(validNeighbors, string(matrix[pos.Row][pos.Col]))
		}
	}

	return validNeighbors
}

func part1(lines []string) int {
	var numbers []int
	var number string
	var hasSymbolNeighbor string

	for row, line := range lines {
		for col, ch := range line + "." {
			if ch != '.' && !is_symbol(string(ch)) {
				for _, rn := range []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} {
					if ch == rn {
						number += string(ch)

						if len(hasSymbolNeighbor) == 0 {
							for _, neighbor := range get_neighbors_part_1(lines, row, col) {
								if is_symbol(neighbor) {
									hasSymbolNeighbor = neighbor
								}
							}
						}
					}
				}
			} else {
				if len(hasSymbolNeighbor) > 0 && len(number) > 0 {
					number_int, _ := strconv.Atoi(number)
					numbers = append(numbers, number_int)
				}
				number = ""
				hasSymbolNeighbor = ""
			}
		}
	}

	return sum_list(numbers)
}

type SymbolList struct {
	symbol string
	coor   Coor
}

type NumberList struct {
	number int
	coors  []Coor
}

func get_neighbors_part_2(matrix []string, row, col int) []Coor {
	rowsLen := len(matrix) - 1
	colsLen := len(matrix[0]) - 1

	if rowsLen == 0 {
		return nil
	}

	positions := []Coor{
		{row - 1, col},     // top
		{row - 1, col + 1}, // top-right
		{row, col + 1},     // right
		{row + 1, col + 1}, // bottom-right
		{row + 1, col},     // bottom
		{row + 1, col - 1}, // bottom-left
		{row, col - 1},     // left
		{row - 1, col - 1}, // top-left
	}

	var validNeighbors []Coor
	for _, pos := range positions {
		if pos.Row >= 0 && pos.Row < rowsLen && pos.Col >= 0 && pos.Col < colsLen {
			validNeighbors = append(validNeighbors, pos)
		}
	}

	return validNeighbors
}

func part2(lines []string) int {
	var number string
	var coors []Coor

	var numbers []NumberList
	var symbols []SymbolList

	for row, line := range lines {
		for col, ch := range line + "." {
			if unicode.IsDigit(ch) {
				number += string(ch)
				coors = append(coors, Coor{row, col})
			} else {
				if len(number) > 0 {
					number_int, _ := strconv.Atoi(number)

					numbers = append(numbers, NumberList{number: number_int, coors: coors})
					number = ""
					coors = []Coor{}
				}
				if ch == '*' {
					symbols = append(symbols, SymbolList{symbol: string(ch), coor: Coor{row, col}})
				}

			}
		}
	}

	var gears_and_neighbors = make(map[int][]int)
	for len(numbers) > 0 {
		number_to_check := numbers[0]
		numbers = numbers[1:]

		found := false
		for i, symbol := range symbols {
			for _, coor := range number_to_check.coors {
				for _, neighbor := range get_neighbors_part_2(lines, coor.Row, coor.Col) {
					if found {
						continue
					}

					if neighbor.Row == symbol.coor.Row && neighbor.Col == symbol.coor.Col {
						if len(gears_and_neighbors[i]) == 0 {
							gears_and_neighbors[i] = []int{number_to_check.number}
						} else {
							gears_and_neighbors[i] = append(gears_and_neighbors[i], number_to_check.number)
						}
						found = true
					}
				}
			}
		}
	}

	sum := 0
	for _, numbers := range gears_and_neighbors {
		if len(numbers) == 2 {
			sum += multiply(numbers)

		}
	}

	return sum
}

func main() {
	data, err := os.ReadFile("day_3/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Result for part1: %d\n", part1(lines)) // 550064
	fmt.Printf("Result for part2: %d\n", part2(lines)) // 85010461
}
