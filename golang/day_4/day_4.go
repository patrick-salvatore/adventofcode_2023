package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type FindWinningMatchesPredicate func(map[string]bool, []string)

func map_winning_matches(line string, prefix_re *regexp.Regexp, number_re *regexp.Regexp, predicate FindWinningMatchesPredicate) {
	winning_map := make(map[string]bool)
	prefix := prefix_re.FindStringSubmatch(line)[0]
	line = strings.Trim(line[len(prefix)-1:], " ")

	allCards := strings.Split(line, "|")
	winning_cards := number_re.FindAllStringSubmatch(allCards[0], -1)
	your_cards := number_re.FindAllStringSubmatch(allCards[1], -1)

	for _, winning_card := range winning_cards {
		winning_map[winning_card[0]] = true
	}

	for _, your_card := range your_cards {
		predicate(winning_map, your_card)
	}
}

func part1(lines []string) int {
	sum := 0

	prefix_re := regexp.MustCompile(`Card\s+(\d+):\s+`)
	number_re := regexp.MustCompile(`(\d+)`)

	for _, line := range lines {
		total := 0
		map_winning_matches(line, prefix_re, number_re, func(m map[string]bool, s []string) {
			if m[s[0]] {
				if total == 0 {
					total = 1
				} else {
					total <<= 1
				}
			}
		})
		sum += total
	}

	return sum
}

func part2(lines []string) int {
	prefix_re := regexp.MustCompile(`Card\s+(\d+):\s+`)
	number_re := regexp.MustCompile(`(\d+)`)

	cards := make(map[int]int)

	sum := 0
	for i, line := range lines {
		total := 0

		map_winning_matches(line, prefix_re, number_re, func(m map[string]bool, s []string) {
			if m[s[0]] {
				total += 1
			}
		})

		if cards[i] == 0 {
			cards[i] = 1
		}

		for j := i + 1; j < i+total+1; j++ {
			if cards[j] > 0 {
				cards[j] = cards[j] + cards[i]
			} else {
				cards[j] = 1 + cards[i]
			}
		}
	}

	for _, v := range cards {
		sum += v
	}

	return sum
}

func main() {
	data, err := os.ReadFile("day_4/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
