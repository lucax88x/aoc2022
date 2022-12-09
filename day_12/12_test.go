package day12

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput12(t *testing.T) {
	ex := ``

	shared.HelperTest(Day12_1, ex, 1, t)
	shared.HelperTest(Day12_2, ex, 1, t)
}

func TestRealInput12(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day12_1, content, 1, t)
	shared.HelperTest(Day12_2, content, 1, t)
}
