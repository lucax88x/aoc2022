package main

import (
	"fmt"
	"os"
	"testing"
)

func HelperTest(sut func(string) int, input string, expected int, t *testing.T) {
	result := sut(input)
	if result != expected {
		t.Errorf("should be %d but is %d", expected, result)
	}
}

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

	HelperTest(Part1_1, ex, 24000, t)
	HelperTest(Part1_2, ex, 45000, t)
}

func TestRealInput(t *testing.T) {
	content, err := os.ReadFile("1_input")

	if err != nil {
		fmt.Print(err)
	}

	HelperTest(Part1_1, string(content), 67622, t)
	HelperTest(Part1_2, string(content), 201491, t)
}
