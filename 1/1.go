package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1_1(str string) int {
	lines := strings.Split(str, "\n")

	largestSum := 0
	currentSum := 0
	for _, line := range lines {

		is_empty_line := len(line) == 0

		if is_empty_line {
			if currentSum > largestSum {
				largestSum = currentSum
			}

			currentSum = 0

			continue
		}

		currentSum += parseLine(line)
	}

	return largestSum
}

func Part1_2(str string) int {
	max := math.MaxInt16
	lines := strings.Split(str, "\n")

	largestThreeSums := []int{max, max, max}

	currentSum := 0
	for line_index, line := range lines {

		is_last_line := line_index == len(lines)-1
		is_empty_line := len(line) == 0

		if is_empty_line || is_last_line {

			if is_last_line && !is_empty_line {
				currentSum = parseLine(line)
			}

			for i := 0; i < len(largestThreeSums); i++ {
				if largestThreeSums[i] == max || currentSum > largestThreeSums[i] {
					largestThreeSums[i] = currentSum
					break
				}
			}

			sort.Ints(largestThreeSums)

			currentSum = 0

			continue
		}

		currentSum += parseLine(line)
	}

	finalSum := 0
	for _, sum := range largestThreeSums {
		finalSum += sum
	}

	return finalSum
}

func parseLine(line string) int {
	parsed, parseError := strconv.Atoi(line)

	if parseError != nil {
		panic(parseError)
	}

	return parsed
}
