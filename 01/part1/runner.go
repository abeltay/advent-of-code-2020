package aoc

import (
	"log"
	"sort"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	arr := make([]int, len(data))
	for i := range data {
		num, err := strconv.Atoi(data[i])
		if err != nil {
			log.Fatal(err)
		}
		arr[i] = num
	}
	sort.Ints(arr)
	for l := range arr {
		h := len(arr) - 1
		for l < h {
			switch {
			case arr[l]+arr[h] == 2020:
				return arr[l] * arr[h]
			case arr[l]+arr[h] < 2020:
				l++
			default:
				h--
			}
		}
	}
	return 0
}
