package day8

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput8(t *testing.T) {
	ex := ``

	shared.HelperTest(Day8_1, ex, 1, t)
	shared.HelperTest(Day8_2, ex, 1, t)
}

func TestRealInput8(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day8_1, content, 1, t)
	shared.HelperTest(Day8_2, content, 1, t)
}
