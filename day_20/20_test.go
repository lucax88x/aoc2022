package day20

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput20(t *testing.T) {
	ex := ``

	shared.HelperTest(Day20_1, ex, 1, t)
	shared.HelperTest(Day20_2, ex, 1, t)
}

func TestRealInput20(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day20_1, content, 1, t)
	shared.HelperTest(Day20_2, content, 1, t)
}
