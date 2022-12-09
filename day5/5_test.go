package day5

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput5(t *testing.T) {
	ex := ``

	shared.HelperTest(Day5_1, ex, 1, t)
	shared.HelperTest(Day5_2, ex, 1, t)
}

func TestRealInput5(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day5_1, content, 1, t)
	shared.HelperTest(Day5_2, content, 1, t)
}
