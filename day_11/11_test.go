package day11

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput11(t *testing.T) {
	ex := ``

	shared.HelperTest(Day11_1, ex, 1, t)
	shared.HelperTest(Day11_2, ex, 1, t)
}

func TestRealInput11(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day11_1, content, 1, t)
	shared.HelperTest(Day11_2, content, 1, t)
}
