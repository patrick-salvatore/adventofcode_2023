package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func create_tree(lines []string) map[string][]string {
	var nodes = make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		if len(nodes[parts[0]]) > 0 {
			panic(fmt.Errorf("node %v, already exists in the tree", nodes[parts[0]]))
		}

		pattern := `\(([^,]+),\s*([^)]+)\)`
		matches := regexp.MustCompile(pattern).FindStringSubmatch(parts[1])
		next := []string{matches[1], matches[2]}
		nodes[parts[0]] = next
	}
	return nodes
}

func part1(lines []string) int {
	instructions := strings.Split(lines[0], "")
	tree_parts := strings.Split(lines[1], "\n")
	tree := create_tree(tree_parts)
	curr := "AAA"

	result := 0
	for curr != "ZZZ" {
		if instructions[0] == "L" {
			curr = tree[curr][0]
		} else {
			curr = tree[curr][1]
		}
		instructions = append(instructions[1:], instructions[0])
		result++
	}

	return result
}

func get_all_starting_points(lines []string) []string {
	var points []string

	for _, line := range lines {
		start := strings.Split(line, " = ")[0]

		if start[len(start)-1] == 'A' {
			points = append(points, start)
		}
	}

	return points
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonDivisor(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = result * numbers[i] / gcd(result, numbers[i])
	}

	return result
}

func part2(lines []string) int {
	instructions := strings.Split(lines[0], "")
	tree_parts := strings.Split(lines[1], "\n")
	tree := create_tree(tree_parts)

	points := get_all_starting_points(tree_parts)

	cycles := [][]int{}

	for _, current := range points {
		cycle := []int{}

		current_instructions := instructions
		step_count := 0
		first_z := ""

		for {
			for step_count == 0 || string(current[len(current)-1]) != "Z" {
				step_count += 1

				if current_instructions[0] == "L" {
					current = tree[current][0]
				} else {
					current = tree[current][1]
				}
				current_instructions = append(current_instructions[1:], current_instructions[0])
			}

			cycle = append(cycle, step_count)

			if first_z == "" {
				first_z = current
				step_count = 0
			} else if current == first_z {
				break
			}
		}
		cycles = append(cycles, cycle)
	}

	nums := []int{}

	for _, c := range cycles {
		nums = append(nums, c[0])
	}

	return leastCommonDivisor(nums)
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
