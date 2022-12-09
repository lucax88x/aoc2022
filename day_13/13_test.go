package day13

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput13(t *testing.T) {
	ex := ``

	shared.HelperTest(Day13_1, ex, 1, t)
	shared.HelperTest(Day13_2, ex, 1, t)
}

func TestRealInput13(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day13_1, content, 1, t)
	shared.HelperTest(Day13_2, content, 1, t)
}
