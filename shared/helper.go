package shared

import (
	"os"
	"testing"
)

func HelperTest(sut func(string) int, input string, expected int, t *testing.T) {
	result := sut(input)
	if result != expected {
		t.Errorf("%s: should be %d but is %d", input, expected, result)
	}
}

func ReadFile(file string) string {
	content, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return string(content)
}
