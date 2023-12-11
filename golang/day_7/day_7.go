package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func _map_iota_to_kind(kind int) string {
	switch kind {
	case HighCard:
		return "HighCard"
	case Pair:
		return "Pair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	}

	panic("UNREACHABLE")
}

type HandsMap map[rune]int

type HandsMapToBid struct {
	h string
	m HandsMap
	b int
	w int
}

func parse_line_part_1(line string) HandsMapToBid {
	var m = make(HandsMap)

	parts := strings.Split(line, " ")
	hand := parts[0]
	bid, _ := strconv.Atoi(parts[1])

	has_four_of_kind := false
	has_three_of_kind := false
	for _, c := range hand {
		m[c] += 1

		if m[c] == 4 {
			has_four_of_kind = true
		} else if m[c] == 3 {
			has_three_of_kind = true
		}
	}

	weight := -1
	switch len(m) {
	case 1:
		weight = FiveOfAKind
	case 2:
		if has_four_of_kind {
			weight = FourOfAKind
		} else {
			weight = FullHouse
		}
	case 3:
		if has_three_of_kind {
			weight = ThreeOfAKind
		} else {
			weight = TwoPair
		}
	case 4:
		weight = Pair
	default:
		weight = HighCard
	}

	return HandsMapToBid{
		h: hand,
		m: m,
		b: bid,
		w: weight,
	}
}

func part1(lines []string) int {
	hands := []HandsMapToBid{}

	for _, line := range lines {
		hands = append(hands, parse_line_part_1(line))
	}

	sort.SliceStable(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		if a.w > b.w {
			return true
		}

		if a.w < b.w {
			return false
		}

		cardValues := map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

		for i := range a.h {
			if cardValues[a.h[i]] == cardValues[b.h[i]] {
				continue
			}
			return cardValues[a.h[i]] > cardValues[b.h[i]]
		}
		return false
	})

	out := 0
	for i := len(hands) - 1; i >= 0; i-- {
		hand := hands[i]
		out += (hand.b * (len(hands) - i))
	}

	return out
}

func map_values_to_slice(_map HandsMap) []int {
	var values []int

	for _, v := range _map {
		values = append(values, v)
	}

	return values
}

func parse_line_part_2(line string) HandsMapToBid {
	var m = make(HandsMap)

	parts := strings.Split(line, " ")
	hand := parts[0]
	bid, _ := strconv.Atoi(parts[1])

	has_four_of_kind := false
	has_three_of_kind := false

	_ = has_four_of_kind || has_three_of_kind
	number_of_jockers := 0
	for _, c := range hand {
		if c == 'J' {
			number_of_jockers += 1
			continue
		}

		m[c] += 1
		if m[c] == 4 {
			has_four_of_kind = true
		} else if m[c] == 3 {
			has_three_of_kind = true
		}
	}

	weight := 0
	switch len(m) {
	case 1:
		weight = FiveOfAKind
	case 2:
		if has_four_of_kind {
			weight = FourOfAKind
		} else {
			weight = FullHouse
		}
	case 3:
		if has_three_of_kind {
			weight = ThreeOfAKind
		} else {
			weight = TwoPair
		}
	case 4:
		weight = Pair
	default:
		weight = HighCard
	}

	if number_of_jockers > 0 {
		// fmt.Printf("hand: %v, weight:%v, ", hand, weight)
		// fmt.Printf("hand: %v, ", hand)
	}

	if number_of_jockers == 5 {
		weight = FiveOfAKind
	}
	if number_of_jockers == 4 {
		weight = FiveOfAKind
	}
	if number_of_jockers == 3 {
		if len(m) == 2 {
			weight = FourOfAKind
		} else if len(m) == 1 {
			weight = FiveOfAKind
		}
	}
	if number_of_jockers == 2 {
		if len(m) == 3 {
			weight = ThreeOfAKind
		} else if len(m) == 2 {
			values := map_values_to_slice(m)
			if values[0] == 2 || values[1] == 2 {
				weight = FourOfAKind
			} else {
				weight = FullHouse
			}
		}
	}
	if number_of_jockers == 1 {
		if len(m) == 4 {
			weight = Pair
		} else if len(m) == 3 {
			weight = ThreeOfAKind
		} else if len(m) == 2 {
			values := map_values_to_slice(m)
			if values[0] == 3 || values[1] == 3 {
				weight = FourOfAKind
			} else {
				weight = FullHouse
			}
		} else if len(m) == 1 {
			weight = FiveOfAKind
		}
	}

	if number_of_jockers > 0 {
		fmt.Printf("hand: %v, best_kind: %v\n", hand, _map_iota_to_kind(weight))
	}

	return HandsMapToBid{
		h: hand,
		m: m,
		b: bid,
		w: weight,
	}
}

func part2(lines []string) int {
	hands := []HandsMapToBid{}

	for _, line := range lines {
		hands = append(hands, parse_line_part_2(line))
	}

	sort.SliceStable(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		if a.w > b.w {
			return true
		}

		if a.w < b.w {
			return false
		}

		cardValues := map[byte]int{'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'Q': 12, 'K': 13, 'A': 14}

		for i := range a.h {
			if cardValues[a.h[i]] == cardValues[b.h[i]] {
				continue
			}
			return cardValues[a.h[i]] > cardValues[b.h[i]]
		}

		return false
	})

	out := 0
	for i := len(hands) - 1; i >= 0; i-- {
		hand := hands[i]
		out += (hand.b * (len(hands) - i))
	}

	return out
}

func main() {
	data, err := os.ReadFile("day_7/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	fmt.Printf("Result for part1: %d\n", part1(lines))
	fmt.Printf("Result for part2: %d\n", part2(lines))
}
