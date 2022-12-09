package day19

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput19(t *testing.T) {
	ex := ``

	shared.HelperTest(Day19_1, ex, 1, t)
	shared.HelperTest(Day19_2, ex, 1, t)
}

func TestRealInput19(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day19_1, content, 1, t)
	shared.HelperTest(Day19_2, content, 1, t)
}
