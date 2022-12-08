package day2

import (
	"fmt"
	"strings"
)

type HandShape int
type Outcome int

const (
	Rock    HandShape = 1
	Paper             = 2
	Scissor           = 3
)

var (
	handShapeMap = map[string]HandShape{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	}
)

const (
	Victory Outcome = 1
	Draw            = 2
	Loss            = 3
)

type Hand struct {
	shape HandShape
}

func NewHand(char string) *Hand {
	hand := new(Hand)

	handShape, ok := handShapeMap[strings.ToUpper(char)]

	if !ok {
		panic(fmt.Sprintf("Not a valid hand %s", char))
	}

	hand.shape = handShape

	return hand
}

func (hand *Hand) match(other *Hand) Outcome {
	if hand.shape == other.shape {
		return Draw
	}

	if hand.shape == Rock {
		switch other.shape {
		case Scissor:
			return Victory
		case Paper:
			return Loss
		}
	}

	if hand.shape == Paper {
		switch other.shape {
		case Rock:
			return Victory
		case Scissor:
			return Loss
		}
	}

	if hand.shape == Scissor {
		switch other.shape {
		case Paper:
			return Victory
		case Rock:
			return Loss
		}
	}

	panic(fmt.Sprintf("Not a valid outcome with %v %v", hand.shape, other.shape))
}

func (hand *Hand) calculate_shape_points() int {

	switch hand.shape {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	}

	panic(fmt.Sprintf("Not a valid shape with %v", hand.shape))
}

func (outcome Outcome) calculate() int {

	switch outcome {
	case Victory:
		return 6
	case Draw:
		return 3
	}

	return 0
}

func Part2_1(str string) int {
	lines := strings.Split(str, "\n")

	points := 0
	for _, line := range lines {
		if len(line) >= 3 {
			player := NewHand(string(line[0]))
			opponent := NewHand(string(line[2]))

			points += player.calculate_shape_points()
			// fmt.Printf("adding shape %v", player.calculate_shape_points())
			println()

			outcome := player.match(opponent)

			points += outcome.calculate()
			// fmt.Printf("adding outcome %v", outcome.calculate())
		}
	}

	return points
}
