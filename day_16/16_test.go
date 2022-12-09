package day16

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput16(t *testing.T) {
	ex := ``

	shared.HelperTest(Day16_1, ex, 1, t)
	shared.HelperTest(Day16_2, ex, 1, t)
}

func TestRealInput16(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day16_1, content, 1, t)
	shared.HelperTest(Day16_2, content, 1, t)
}
