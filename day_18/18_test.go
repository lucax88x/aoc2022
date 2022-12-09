package day18

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput18(t *testing.T) {
	ex := ``

	shared.HelperTest(Day18_1, ex, 1, t)
	shared.HelperTest(Day18_2, ex, 1, t)
}

func TestRealInput18(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day18_1, content, 1, t)
	shared.HelperTest(Day18_2, content, 1, t)
}
