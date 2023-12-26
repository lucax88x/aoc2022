package dayX

import (
	"aoc2023/shared"
	"testing"
)

func TestExampleInput(t *testing.T) {
	ex1 := ``
	// ex2 := ``

	shared.HelperTest(DayX1 ex1, 0, t)
	// shared.HelperTest(DayX_2, ex2, 0, t)
}

func TestRealInput(t *testing.T) {
	content := shared.ReadFile("input")

	shared.HelperTest(DayX1 content, 0, t)
	// shared.HelperTest(DayX_2, content, 0, t)
}
