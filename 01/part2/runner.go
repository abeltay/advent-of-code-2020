package aoc

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// Runner runs the algorithm to get the answer
func Runner(filename string) int {
	arr := parseFile(filename)

	sort.Ints(arr)
	for i := range arr {
		l := i + 1
		h := len(arr) - 1
		for l < h {
			switch {
			case arr[i]+arr[l]+arr[h] == 2020:
				return arr[i] * arr[l] * arr[h]
			case arr[i]+arr[l]+arr[h] < 2020:
				l++
			default:
				h--
			}
		}
	}
	return 0
}

func parseFile(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []int
	for {
		var num int
		_, err := fmt.Fscanln(f, &num)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, num)
	}
	return arr
}
