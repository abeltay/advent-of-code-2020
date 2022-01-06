package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 286
		input1, input2 := ParseFile("../testdata/example.txt")
		got := Runner(input1, input2)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("full input", func(t *testing.T) {
		input1, input2 := ParseFile("../testdata/input.txt")
		got := Runner(input1, input2)
		fmt.Println(got, "is the answer")
	})
}
