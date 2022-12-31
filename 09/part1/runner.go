package aoc

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(arr []int, spread int) int {
	for i := spread; i < len(arr); i++ {
		if !twoSum(arr, i, spread) {
			return arr[i]
		}
	}
	return 0
}

func twoSum(arr []int, cur, spread int) bool {
	for k := cur - spread; k < cur; k++ {
		for l := k + 1; l < cur; l++ {
			if arr[cur] == arr[k]+arr[l] {
				return true
			}
		}
	}
	return false
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []int {
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

// type line struct {
// 	first  int
// 	second int
// 	letter byte
// 	text   string
// }
