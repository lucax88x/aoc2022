package day7

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput7(t *testing.T) {
	ex := ``

	shared.HelperTest(Day7_1, ex, 1, t)
	shared.HelperTest(Day7_2, ex, 1, t)
}

func TestRealInput7(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day7_1, content, 1, t)
	shared.HelperTest(Day7_2, content, 1, t)
}
