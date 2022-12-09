package day17

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput17(t *testing.T) {
	ex := ``

	shared.HelperTest(Day17_1, ex, 1, t)
	shared.HelperTest(Day17_2, ex, 1, t)
}

func TestRealInput17(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day17_1, content, 1, t)
	shared.HelperTest(Day17_2, content, 1, t)
}
