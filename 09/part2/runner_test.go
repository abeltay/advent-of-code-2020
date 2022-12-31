package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 62
		input := ParseFile("../testdata/example.txt")
		got := Runner(input, 5)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("full input", func(t *testing.T) {
		input := ParseFile("../testdata/input.txt")
		got := Runner(input, 25)
		fmt.Println(got, "is the answer")
	})
}
