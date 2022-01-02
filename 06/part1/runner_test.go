package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 11
		input := ParseFile("../testdata/example.txt")
		got := Runner(input)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("full input", func(t *testing.T) {
		input := ParseFile("../testdata/input.txt")
		got := Runner(input)
		fmt.Println(got, "is the answer")
	})
}
