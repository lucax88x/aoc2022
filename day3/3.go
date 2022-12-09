package day3

import (
	"strings"
	"golang.org/x/exp/slices"
)

var Points = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func Day3_1(str string) int {
	lines := strings.Split(str, "\n")

	points := 0
	for _, line := range lines {

		line_length := len(line)
		half_index := line_length / 2

		left_part := line[0:half_index]
		right_part := line[half_index:line_length]

		pairs := []rune{}

		for _, left_char := range left_part {
			for _, right_char := range right_part {
				if left_char == right_char {
					if !slices.Contains(pairs, left_char) {
						pairs = append(pairs, left_char)
						continue
					}
				}
			}
		}

		points = calculatePoints(pairs, points)
	}

	return points
}

func calculatePoints(pairs []rune, points int) int {
	for _, pair := range pairs {
		points += Points[pair]
	}
	return points
}

func Day3_2(str string) int {
	lines := strings.Split(str, "\n")

	points := 0
	for index := 0; index+2 < len(lines); index += 3 {
		first_group := lines[index]
		second_group := lines[index+1]
		third_group := lines[index+2]

		pairs := []rune{}
		for _, first_group_char := range first_group {
			for _, second_group_char := range second_group {
				if first_group_char == second_group_char {
					for _, third_group_char := range third_group {
						if first_group_char == third_group_char {
							if !slices.Contains(pairs, first_group_char) {
								pairs = append(pairs, first_group_char)
								continue
							}
						}
					}
				}
			}
		}

		points = calculatePoints(pairs, points)
	}

	return points
}
