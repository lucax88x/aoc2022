package day1

import (
	"aoc2023/shared"
	"strconv"
	"strings"
	"unicode"
)

func Day11(str string) int {
	lines := strings.Split(str, "\n")

	sum := 0
	for _, line := range lines {
		sum += parseLine1(line)
	}
	return sum
}

func Day12(str string) int {
	lines := strings.Split(str, "\n")

	sum := 0
	for _, line := range lines {
		sum += parseLine2(line)
	}
	return sum

}

func parseLine1(line string) int {
	left := '0'
	right := '0'

	for i := 0; i < len(line); i++ {
		char := line[i]
		if unicode.IsDigit(rune(char)) {
			left = rune(char)
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]
		if unicode.IsDigit(rune(char)) {
			right = rune(char)
			break
		}
	}

	parsed, parseError := strconv.Atoi(string(left) + string(right))

	if parseError != nil {
		panic(parseError)
	}

	return parsed
}

func parseLine2(line string) int {
	if len(line) == 0 {
		return 0
	}

	left := '0'
	right := '0'
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	mostLeftIndex := -1
	mostRightIndex := -1
	mostRightDigitLength := 0

	for digitIndex, digit := range digits {
		leftIndex := strings.Index(line, digit)
		rightIndex := strings.LastIndex(line, digit)

		if leftIndex > -1 {
			if mostLeftIndex == -1 || leftIndex < mostLeftIndex {
				mostLeftIndex = leftIndex
				left = rune(digitIndex+1) + '0'
				shared.Debug("found left! %c", left)
			}
		}
		if rightIndex > -1 {
			if mostRightIndex == -1 || rightIndex >= mostRightIndex {
				mostRightIndex = rightIndex
				mostRightDigitLength = len(digit)
				right = rune(digitIndex+1) + '0'
				shared.Debug("found right! %c", right)
			}

		}
	}

	if mostLeftIndex == -1 {
		mostLeftIndex = len(line) - 1
	}
	if mostRightIndex == -1 {
		mostRightIndex = 0
	}

	shared.Debug("starting left until index %d", mostLeftIndex)
	for i := 0; i <= mostLeftIndex; i++ {
		char := line[i]
		if unicode.IsDigit(rune(char)) {
			left = rune(char)
			shared.Debug("found left as digit! %c", char)
			break
		}
	}

	shared.Debug("starting right until index %d", mostRightIndex+mostRightDigitLength)
	for i := len(line) - 1; i >= mostRightIndex+mostRightDigitLength; i-- {
		char := line[i]
		if unicode.IsDigit(rune(char)) {
			right = rune(char)
			shared.Debug("found right as digit! %c", char)
			break
		}
	}

	shared.Debug("left: %c right: %c", left, right)
	println()
	parsed, parseError := strconv.Atoi(string(left) + string(right))

	if parseError != nil {
		panic(parseError)
	}

	return parsed
}
