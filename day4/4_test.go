package day4

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput4(t *testing.T) {
	ex := ``

	shared.HelperTest(Day4_1, ex, 1, t)
	shared.HelperTest(Day4_2, ex, 1, t)
}

func TestRealInput4(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day4_1, content, 1, t)
	shared.HelperTest(Day4_2, content, 1, t)
}
