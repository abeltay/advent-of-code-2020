package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("full input", func(t *testing.T) {
		input := ParseFile("../testdata/input.txt")
		got := Runner(input)
		fmt.Println(got, "is the answer")
	})
}
