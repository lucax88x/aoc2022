package day14

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput14(t *testing.T) {
	ex := ``

	shared.HelperTest(Day14_1, ex, 1, t)
	shared.HelperTest(Day14_2, ex, 1, t)
}

func TestRealInput14(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day14_1, content, 1, t)
	shared.HelperTest(Day14_2, content, 1, t)
}
