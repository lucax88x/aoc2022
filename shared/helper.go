package shared

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func HelperTest(sut func(string) int, input string, expected int, t *testing.T) {
	result := sut(input)
	if result != expected {
		t.Errorf("should be %d but is %d", expected, result)
	}
}

func ReadFile(file string) string {
	content, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func Debug(format string, a ...any) {
	fmt.Printf(format, a...)
	println()
}

func ReadInput() string {
	_, file, _, _ := runtime.Caller(1)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input")

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(input))
}

func ReadInputLines() []string {
	_, file, _, _ := runtime.Caller(1)
	dir := path.Dir(file)
	input, err := os.ReadFile(dir + "/input")

	if err != nil {
		return []string{}
	}

	return strings.Split(strings.TrimSpace(string(input)), "\n")
}

func ToInt(value string) int {
	result, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil {
		panic(err)
	}

	return result
}
