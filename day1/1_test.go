package day1

import (
	"aoc2023/shared"
	"testing"
)

func TestExampleInput(t *testing.T) {
	ex1 := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	ex2 := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	shared.HelperTest(Day11, ex1, 142, t)
	shared.HelperTest(Day12, ex2, 281, t)
}

func TestRealInput(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day11, content, 55123, t)
	shared.HelperTest(Day12, content, 55260, t)
}
