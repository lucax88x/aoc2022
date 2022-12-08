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

var (
	outcomeMap = map[string]Outcome{
		"X": Loss,
		"Y": Draw,
		"Z": Victory,
	}
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

func PredictHand(otherHand *Hand, char string) *Hand {
	expected_outcome := predict_outcome(char)

	hand := new(Hand)
	hand.shape = predict_hand_shape(otherHand, expected_outcome)

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

func predict_hand_shape(other_hand *Hand, expected_outcome Outcome) HandShape {
	if expected_outcome == Draw {
		return other_hand.shape
	}

	if other_hand.shape == Rock {
		switch expected_outcome {
		case Victory:
			return Paper
		case Loss:
			return Scissor
		}
	}

	if other_hand.shape == Paper {
		switch expected_outcome {
		case Victory:
			return Scissor
		case Loss:
			return Rock
		}

	}

	if other_hand.shape == Scissor {
		switch expected_outcome {
		case Victory:
			return Rock
		case Loss:
			return Paper
		}
	}

	panic(fmt.Sprintf("Not a valid predicted hand with %v %v", other_hand.shape, expected_outcome))

}

func predict_outcome(char string) Outcome {
	expected_outcome, ok := outcomeMap[strings.ToUpper(char)]

	if !ok {
		panic(fmt.Sprintf("Not a outcome %s", char))
	}

	return expected_outcome
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

func Day2_1(str string) int {
	lines := strings.Split(str, "\n")

	points := 0
	for _, line := range lines {
		if len(line) >= 3 {
			opponent := NewHand(string(line[0]))
			player := NewHand(string(line[2]))

			points += player.calculate_shape_points()

			outcome := player.match(opponent)

			points += outcome.calculate()
		}
	}

	return points
}

func Day2_2(str string) int {
	lines := strings.Split(str, "\n")

	points := 0
	for _, line := range lines {
		if len(line) >= 3 {
			opponent := NewHand(string(line[0]))
			player := PredictHand(opponent, string(line[2]))

			points += player.calculate_shape_points()

			outcome := player.match(opponent)

			points += outcome.calculate()
		}
	}

	return points
}
