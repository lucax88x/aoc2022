package day1

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput(t *testing.T) {
	ex := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	shared.HelperTest(Day1_1, ex, 24000, t)
	shared.HelperTest(Day1_2, ex, 45000, t)
}

func TestRealInput(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day1_1, content, 67622, t)
	shared.HelperTest(Day1_2, content, 201491, t)
}
