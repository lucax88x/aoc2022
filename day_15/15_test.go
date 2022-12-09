package day15

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput15(t *testing.T) {
	ex := ``

	shared.HelperTest(Day15_1, ex, 1, t)
	shared.HelperTest(Day15_2, ex, 1, t)
}

func TestRealInput15(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day15_1, content, 1, t)
	shared.HelperTest(Day15_2, content, 1, t)
}
