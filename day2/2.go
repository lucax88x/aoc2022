package day2

import (
	"aoc2023/shared"
	"strings"
)

func Day21(str string) int {
	lines := strings.Split(str, "\n")

	sum := 0
	for _, line := range lines {
		sum += parseLine1(line)
	}
	return sum
}

func Day22(str string) int {
	lines := strings.Split(str, "\n")

	sum := 0
	for _, line := range lines {
		sum += parseLine2(line)
	}
	return sum

}

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

func parseLine1(line string) int {
	header := strings.Split(line, ":")

	if len(header) != 2 {
		return 0
	}

	gameNumber := shared.ToInt(strings.Replace(header[0], "Game ", "", 1))
	sets := strings.Split(header[1], ";")

	for _, set := range sets {
		extractedRed := 0
		extractedGreen := 0
		extractedBlue := 0

		rounds := strings.Split(set, ",")
		for _, round := range rounds {
			if strings.Contains(round, "red") {
				extractedRed += shared.ToInt(strings.Replace(round, " red", "", 1))

				if extractedRed > maxRed {
					return 0
				}
			}
			if strings.Contains(round, "green") {
				extractedGreen += shared.ToInt(strings.Replace(round, " green", "", 1))

				if extractedGreen > maxGreen {
					return 0
				}
			}
			if strings.Contains(round, "blue") {
				extractedBlue += shared.ToInt(strings.Replace(round, " blue", "", 1))

				if extractedBlue > maxBlue {
					return 0
				}
			}

		}
	}

	return gameNumber
}

func parseLine2(line string) int {
	header := strings.Split(line, ":")

	if len(header) != 2 {
		return 0
	}

	sets := strings.Split(header[1], ";")

	highestRedFound := 0
	highestGreenFound := 0
	highestBlueFound := 0

	for _, set := range sets {
		rounds := strings.Split(set, ",")
		for _, round := range rounds {
			if strings.Contains(round, "red") {
				extractedRed := shared.ToInt(strings.Replace(round, " red", "", 1))

				if extractedRed > highestRedFound {
					highestRedFound = extractedRed
				}
			}
			if strings.Contains(round, "green") {
				extractedGreen := shared.ToInt(strings.Replace(round, " green", "", 1))

				if extractedGreen > highestGreenFound {
					highestGreenFound = extractedGreen
				}
			}
			if strings.Contains(round, "blue") {
				extractedBlue := shared.ToInt(strings.Replace(round, " blue", "", 1))

				if extractedBlue > highestBlueFound {
					highestBlueFound = extractedBlue
				}
			}

		}
	}

	return highestRedFound * highestBlueFound * highestGreenFound
}
