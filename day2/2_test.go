package day2

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput2(t *testing.T) {
	ex := `A Y
B X
C Z`

	shared.HelperTest(Day2_1, ex, 15, t)
	shared.HelperTest(Day2_2, ex, 12, t)
}

func TestCasesInput2(t *testing.T) {
	shared.HelperTest(Day2_1, "A X", 4, t) // rock vs rock: 1 + 3
	shared.HelperTest(Day2_1, "A Y", 8, t) // rock vs paper: 2 + 6
	shared.HelperTest(Day2_1, "A Z", 3, t) // rock vs scissor: 3 + 0

	shared.HelperTest(Day2_1, "B X", 1, t) // paper vs rock: 1 + 0
	shared.HelperTest(Day2_1, "B Y", 5, t) // paper vs paper: 2 + 3
	shared.HelperTest(Day2_1, "B Z", 9, t) // paper vs scissor: 3 + 6

	shared.HelperTest(Day2_1, "C X", 7, t) // scissor vs rock: 1 + 6
	shared.HelperTest(Day2_1, "C Y", 2, t) // scissor vs paper: 2 + 0
	shared.HelperTest(Day2_1, "C Z", 6, t) // scissor vs scissor: 3 + 3
  
	shared.HelperTest(Day2_2, "A X", 3, t) // rock => scissor loss: 3 + 0
	shared.HelperTest(Day2_2, "A Y", 4, t) // rock => rock draw: 1 + 3
	shared.HelperTest(Day2_2, "A Z", 8, t) // rock => paper vs victory: 2 + 6

	shared.HelperTest(Day2_2, "B X", 1, t) // paper vs loss: 1 + 0
	shared.HelperTest(Day2_2, "B Y", 5, t) // paper vs draw: 2 + 3
	shared.HelperTest(Day2_2, "B Z", 9, t) // paper vs victory: 3 + 6

	shared.HelperTest(Day2_2, "C X", 2, t) // scissor => paper  loss: 2 + 0
	shared.HelperTest(Day2_2, "C Y", 6, t) // scissor => scissor vs draw: 3 + 3
	shared.HelperTest(Day2_2, "C Z", 7, t) // scissor => rock vs victory: 1 + 6
}

func TestRealInput2(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Day2_1, content, 11386, t)
	shared.HelperTest(Day2_2, content, 13600, t)
}
