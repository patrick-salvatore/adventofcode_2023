package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var INT_MAP = map[string]string{
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func is_digit(ch rune) int {
	if unicode.IsDigit(ch) {
		num, _ := strconv.Atoi(string(ch))
		return num
	}

	return -1
}

func find_digits(str string) []int {
	var digits = []int{}

	for _, ch := range str {
		num := is_digit(ch)

		if num > 0 {
			digits = append(digits, num)
		}
	}

	return digits
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {

		digits := find_digits(line)
		first, last := digits[0], digits[len(digits)-1]

		sum += ((first * 10) + (last))
	}

	return sum
}

func find_first_digit(line string) (int, int) {
	for i, ch := range line {
		num := is_digit(ch)
		if num > 0 {
			return i, num
		}
	}

	return -1, -1
}

func find_first_number(line string) (int, int) {
	temp_index := math.MaxInt
	found_number := -1
	for _, word := range []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	} {

		found := strings.Index(line, word)
		if found > -1 && found < temp_index {
			num, _ := strconv.Atoi(string(INT_MAP[word]))
			found_number = num
			temp_index = found
		}

		// fmt.Printf("found: %d, temp_index: %d, word: %s\n", found_number, temp_index, word)
	}
	return temp_index, found_number
}

func find_last_digit(line string) (int, int) {
	for i := len(line) - 1; i > -1; i-- {
		ch := line[i]
		num := is_digit(rune(ch))
		if num > -1 {
			return i, num
		}
	}

	return -1, -1
}

func find_last_number(line string) (int, int) {
	temp_index := 0
	found_number := -1
	for _, word := range []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	} {

		found := strings.LastIndex(line, word)
		if found > temp_index {
			num, _ := strconv.Atoi(string(INT_MAP[word]))
			found_number = num
			temp_index = found
		}
	}

	return temp_index, found_number
}

func part2(lines []string) int {

	sum := 0
	for _, line := range lines {
		firstDigitIndex, firstDigit := find_first_digit(line)
		firstNumberIndex, firstNumber := find_first_number(line)
		lastDigitIndex, lastDigit := find_last_digit(line)
		lastNumberIndex, lastNumber := find_last_number(line)

		first := 0
		last := 0

		// this will never fail but i like to check anyway
		if firstDigitIndex <= firstNumberIndex {
			first = firstDigit
		} else {
			first = firstNumber
		}

		if lastNumber > -1 {
		}
		// this will never fail but i like to check anyway
		if lastDigitIndex >= lastNumberIndex {
			last = lastDigit
		} else {
			last = lastNumber
		}

		sum += (first * 10) + last
	}

	return sum
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
