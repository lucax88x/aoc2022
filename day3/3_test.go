package day3

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput2(t *testing.T) {
	ex := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	shared.HelperTest(Day3_1, ex, 157, t)
	shared.HelperTest(Day3_2, ex, 70, t)
}

func TestRealInput2(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day3_1, content, 8072, t)
	shared.HelperTest(Day3_2, content, 2567, t)
}
