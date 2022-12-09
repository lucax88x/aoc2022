package day9

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput9(t *testing.T) {
	ex := ``

	shared.HelperTest(Day9_1, ex, 1, t)
	shared.HelperTest(Day9_2, ex, 1, t)
}

func TestRealInput9(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day9_1, content, 1, t)
	shared.HelperTest(Day9_2, content, 1, t)
}
