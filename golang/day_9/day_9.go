package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func everything_0(n []int) bool {
	for _, i := range n {
		if i != 0 {
			return false
		}
	}

	return true
}

func get_pairs(input []int) [][]int {
	var result [][]int

	for i := 0; i < len(input)-1; i++ {
		pair := []int{input[i], input[i+1]}
		result = append(result, pair)
	}

	return result
}

func do_recur_part_1(ints []int) int {
	if everything_0(ints) {
		return 0
	}

	next := []int{}
	for _, pair := range get_pairs(ints) {
		next = append(next, pair[1]-pair[0])
	}

	diff := do_recur_part_1(next)
	return ints[len(ints)-1] + diff

}

func string_list_to_int_list(list []string) []int {
	nums := []int{}
	for _, i := range list {
		n, _ := strconv.Atoi(i)
		nums = append(nums, n)
	}

	return nums
}

func part1(lines []string) int {
	result := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := string_list_to_int_list(strings.Split(line, " "))
		result += do_recur_part_1(numbers)
	}

	return result
}

func do_recur_part_2(ints []int) int {
	if everything_0(ints) {
		return 0
	}

	next := []int{}
	for _, pair := range get_pairs(ints) {
		next = append(next, pair[1]-pair[0])
	}

	diff := do_recur_part_2(next)
	return ints[0] - diff

}

func part2(lines []string) int {
	result := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := string_list_to_int_list(strings.Split(line, " "))
		result += do_recur_part_2(numbers)
	}

	return result
}

func main() {
	data, err := os.ReadFile("day_9/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
