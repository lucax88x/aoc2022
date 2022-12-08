package day2

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput2(t *testing.T) {
	ex := `A Y
B X
C Z`

	shared.HelperTest(Part2_1, ex, 15, t)
	// HelperTest(Part1_2, ex, 45000, t)
}

func TestRealInput2(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Part2_1, content, 0, t)
	// HelperTest(Part2_2, content, 201491, t)
}
