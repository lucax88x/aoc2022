package day4

import (
	"strconv"
	"strings"
)

func Day4_1(str string) int {
	return parse_pairs(str, func(left_pair_left_bound int,
		left_pair_right_bound int,
		right_pair_left_bound int,
		right_pair_right_bound int) int {
		if left_pair_left_bound <= right_pair_left_bound && left_pair_right_bound >= right_pair_right_bound {
			return 1
		}

		if right_pair_left_bound <= left_pair_left_bound && right_pair_right_bound >= left_pair_right_bound {
			return 1
		}

		return 0
	})
}

func Day4_2(str string) int {
	return parse_pairs(str, func(left_pair_left_bound int,
		left_pair_right_bound int,
		right_pair_left_bound int,
		right_pair_right_bound int) int {

		left_length := left_pair_right_bound - left_pair_left_bound
		right_length := right_pair_left_bound - right_pair_right_bound

		if left_length != right_length {

		}

		return 0
	})
}

func parseNumber(line string) int {
	parsed, parseError := strconv.Atoi(line)

	if parseError != nil {
		panic(parseError)
	}

	return parsed
}

func parse_pairs(str string, callback func(
	left_pair_left_bound int,
	left_pair_right_bound int,
	right_pair_left_bound int,
	right_pair_right_bound int) int) int {
	lines := strings.Split(str, "\n")

	found := 0

	for _, line := range lines {
		pairs := strings.Split(line, ",")

		if len(pairs) != 2 {
			continue
		}

		left_pair := strings.Split(pairs[0], "-")
		right_pair := strings.Split(pairs[1], "-")

		if len(left_pair) != 2 || len(right_pair) != 2 {
			continue
		}

		left_pair_left_bound := parseNumber(left_pair[0])
		left_pair_right_bound := parseNumber(left_pair[1])

		right_pair_left_bound := parseNumber(right_pair[0])
		right_pair_right_bound := parseNumber(right_pair[1])

		found += callback(left_pair_left_bound, left_pair_right_bound, right_pair_left_bound, right_pair_right_bound)
	}

	return found
}
