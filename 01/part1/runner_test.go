package aoc

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		want := 514579
		got := Runner("../example.txt")
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("full input", func(t *testing.T) {
		got := Runner("../input.txt")
		fmt.Println(got, "is the answer")
	})
}
