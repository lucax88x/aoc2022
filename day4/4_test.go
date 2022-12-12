package day4

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput4(t *testing.T) {
	ex := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	shared.HelperTest(Day4_1, ex, 2, t)
	shared.HelperTest(Day4_2, ex, 4, t)
}

func TestRealInput4(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day4_1, content, 444, t)
	shared.HelperTest(Day4_2, content, 1, t)
}
