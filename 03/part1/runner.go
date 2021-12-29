package aoc

import (
	"bufio"
	"log"
	"os"
)

const (
	incRight = 3
	incDown  = 1
)

// Runner runs the algorithm to get the answer
func Runner(filename string) int {
	arr := parseFile(filename)
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }

	var ans int
	for i := 1; i < len(arr); i++ {
		right := incRight * i
		right = right % len(arr[0])
		if arr[i][right] {
			ans++
		}
	}
	return ans
}

func parseFile(filename string) [][]bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr [][]bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		barr := make([]bool, len(t))
		for i := range t {
			if t[i] == '#' {
				barr[i] = true
			}
		}
		arr = append(arr, barr)
	}
	return arr
}
