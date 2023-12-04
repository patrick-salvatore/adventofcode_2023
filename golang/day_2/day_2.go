package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func do_part1(line string) int {
	//						     R   G	 B
	var COLORS = []int{12, 13, 14}
	gameRe := regexp.MustCompile(`Game\s*(\d+)`)
	colorRe := regexp.MustCompile(`(\d+)\s*(red|green|blue)`)

	gameMatches := gameRe.FindAllStringSubmatch(line, -1)
	gameId, _ := strconv.Atoi(string(strings.Split(gameMatches[0][0], " ")[1]))
	game := strings.Trim(strings.Split(line, ":")[1], " ")
	subsets := strings.Split(game, ";")

	for _, subset := range subsets {
		subset = strings.Trim(subset, " ")
		colorMatches := colorRe.FindAllStringSubmatch(subset, -1)

		for _, colorMatch := range colorMatches {
			found := strings.Split(strings.Trim(colorMatch[0], " "), " ")
			num, _ := strconv.Atoi(string(found[0]))

			color := found[1]
			if color == "red" && num > COLORS[0] {
				return 0
			}
			if color == "green" && num > COLORS[1] {
				return 0
			}
			if color == "blue" && num > COLORS[2] {
				return 0
			}
		}
	}

	return gameId
}

func part1(lines []string) int {
	var SUM = 0

	for _, line := range lines {
		SUM += do_part1(line)
	}

	return SUM
}

func multiply(list []int) int {
	total := 1

	for _, num := range list {
		total *= num
	}

	return total
}

func do_part2(line string) int {
	colorRe := regexp.MustCompile(`(\d+)\s*(red|green|blue)`)
	game := strings.Trim(strings.Split(line, ":")[1], " ")
	subsets := strings.Split(game, ";")
	acc := []int{1, 1, 1}
	for _, subset := range subsets {
		subset = strings.Trim(subset, " ")
		colorMatches := colorRe.FindAllStringSubmatch(subset, -1)

		for _, colorMatch := range colorMatches {
			found := strings.Split(strings.Trim(colorMatch[0], " "), " ")
			num, _ := strconv.Atoi(string(found[0]))
			color := found[1]

			if color == "red" && num > acc[0] {
				acc[0] = num
			}
			if color == "green" && num > acc[1] {
				acc[1] = num
			}
			if color == "blue" && num > acc[2] {
				acc[2] = num
			}
		}
	}

	return multiply(acc)
}

func part2(lines []string) int {
	SUM := 0
	for _, line := range lines {
		SUM += do_part2(line)
	}

	return SUM
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
