package day2

import (
	"aoc2022/shared"
	"testing"
)

func TestExampleInput2(t *testing.T) {

	shared.HelperTest(Part2_1, `A Y
B X
C Z`, 15, t)
	// HelperTest(Part1_2, ex, 45000, t)
}

func TestCasesInput2(t *testing.T) {
  shared.HelperTest(Part2_1, "A X", 4, t) // rock vs rock: 1 + 3
  shared.HelperTest(Part2_1, "A Y", 8, t) // rock vs paper: 2 + 6
  shared.HelperTest(Part2_1, "A Z", 3, t) // rock vs scissor: 3 + 0
  
  shared.HelperTest(Part2_1, "B X", 1, t) // paper vs rock: 1 + 0
  shared.HelperTest(Part2_1, "B Y", 5, t) // paper vs paper: 2 + 3
  shared.HelperTest(Part2_1, "B Z", 9, t) // paper vs scissor: 3 + 6
  
  shared.HelperTest(Part2_1, "C X", 7, t) // scissor vs rock: 1 + 6
  shared.HelperTest(Part2_1, "C Y", 2, t) // scissor vs paper: 2 + 0
  shared.HelperTest(Part2_1, "C Z", 6, t) // scissor vs scissor: 3 + 3
}

func TestRealInput2(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(Part2_1, content, 11386, t)
	// HelperTest(Part2_2, content, 201491, t)
}
