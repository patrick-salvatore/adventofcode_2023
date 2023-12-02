package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {

}
func part2() {

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
