package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func min(numbers []int) int {
	if len(numbers) == 0 {
		// Handle the case where the list is empty
		return math.MaxInt
	}

	minValue := numbers[0]

	for _, num := range numbers {
		if num < minValue {
			minValue = num
		}
	}

	return minValue
}

func parse_seed_line(line string) []int {
	seeds := []int{}

	nums := strings.Split(strings.Trim(strings.Split(line, ": ")[1], " "), " ")

	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		seeds = append(seeds, n)
	}

	return seeds
}

type Range struct {
	a, b, c int
}

func string_to_range(line string) Range {
	parts := strings.Split(strings.Trim(line, " "), " ")

	if len(parts) != 3 {
		fmt.Println("line ", line)
		fmt.Println("parts ", parts)

		os.Exit(1)
	}

	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])

	return Range{
		a, b, c,
	}
}

func part1(lines []string) int {
	seeds := parse_seed_line(lines[0])

	for _, block := range lines[1:] {
		lines := strings.Split(block, "\n")[1:]

		ranges := []Range{}
		for _, line := range lines {
			ranges = append(ranges, string_to_range(line))
		}

		new_seeds := []int{}

		for _, seed := range seeds {
			found := false
			for _, r := range ranges {
				if r.b <= seed && seed < r.b+r.c {

					new := (seed - r.b) + r.a
					new_seeds = append(new_seeds, new)
					found = true
					break
				}
			}

			if !found {
				new_seeds = append(new_seeds, seed)
			}
		}

		seeds = new_seeds
	}

	return min(seeds)
}

func part2(lines []string) int {
	seeds := parse_seed_line(lines[0])

	for _, block := range lines[1:] {
		lines := strings.Split(block, "\n")[1:]

		ranges := []Range{}
		for _, line := range lines {
			ranges = append(ranges, string_to_range(line))
		}

		new_seeds := []int{}
		for _, seed := range seeds {
			found := false
			for _, r := range ranges {
				if r.b <= seed && seed < r.b+r.c {
					new := (seed - r.b) + r.a
					new_seeds = append(new_seeds, new)
					found = true
					break
				}
			}

			if !found {
				new_seeds = append(new_seeds, seed)
			}
		}

		seeds = new_seeds
	}

	return min(seeds)
}

func main() {
	data, err := os.ReadFile("day_5/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
