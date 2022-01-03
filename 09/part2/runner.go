package aoc

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(arr []int, spread int) int {
	var num int
	for i := spread; i < len(arr); i++ {
		if !twoSum(arr, i, spread) {
			num = arr[i]
			break
		}
	}
	var p1, p2 int
	sum := arr[p1]
	for {
		switch {
		case sum == num:
			min, max := arr[p1], arr[p1]
			for i := p1 + 1; i <= p2; i++ {
				if arr[i] < min {
					min = arr[i]
				}
				if arr[i] > max {
					max = arr[i]
				}
			}
			return min + max
		case sum < num:
			p2++
			sum += arr[p2]
		case sum > num:
			sum -= arr[p1]
			p1++
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
