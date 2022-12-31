package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 2
		file := ParseFile("../testdata/example.txt")
		got := Runner(file)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("full input", func(t *testing.T) {
		file := ParseFile("../testdata/input.txt")
		got := Runner(file)
		fmt.Println(got, "is the answer")
	})
}
