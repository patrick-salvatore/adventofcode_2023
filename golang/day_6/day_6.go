package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_values_part_1(line string) []int {
	fields := strings.Fields(line)

	var values []int

	for _, valueStr := range fields[1:] {
		value, err := strconv.Atoi(strings.TrimSpace(valueStr))
		if err == nil {
			values = append(values, value)
		}
	}

	return values
}

func get_values_part_2(line string) []int {
	fields := strings.Fields(line)

	var values []int

	value, err := strconv.Atoi(strings.Join(fields[1:], ""))
	if err == nil {
		values = append(values, value)
	}

	return values
}

func part1(lines []string) int {
	timeValues := get_values_part_1(lines[0])
	distanceValues := get_values_part_1(lines[1])
	out := 1

	for i, time := range timeValues {
		opts := []int{}

		distance := distanceValues[i]
		ticker := 1

		for ticker < time {
			time_left := time - ticker

			if time_left < time && ((time_left * ticker) > distance) {
				opts = append(opts, ticker)
			}

			ticker++
		}

		if len(opts) > 0 {
			out *= len(opts)
		}

	}

	return out
}

func part2(lines []string) int {
	timeValues := get_values_part_2(lines[0])
	distanceValues := get_values_part_2(lines[1])
	out := 1

	for i, time := range timeValues {
		opts := []int{}

		distance := distanceValues[i]
		ticker := 1

		for ticker < time {
			time_left := time - ticker

			if time_left < time && ((time_left * ticker) > distance) {
				opts = append(opts, ticker)
			}

			ticker++
		}

		if len(opts) > 0 {
			out *= len(opts)
		}

	}

	return out
}

func main() {
	data, err := os.ReadFile("day_6/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
