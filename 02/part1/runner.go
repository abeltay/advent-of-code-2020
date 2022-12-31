package aoc

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(filename string) int {
	arr := parseFile(filename)
	var ans int
	for i := range arr {
		var count int
		for j := range arr[i].text {
			if arr[i].text[j] == arr[i].letter {
				count++
			}
		}
		if arr[i].first <= count && count <= arr[i].second {
			ans++
		}
	}
	return ans
}

func parseFile(filename string) []line {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []line
	for {
		var l line
		_, err := fmt.Fscanf(f, "%d-%d %c: %s", &l.first, &l.second, &l.letter, &l.text)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, l)
	}
	return arr
}

type line struct {
	first  int
	second int
	letter byte
	text   string
}
