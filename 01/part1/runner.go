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
	for i := range arr {
		num := 2020 - arr[i]
		loc := sort.SearchInts(arr, num)
		if loc < len(arr) && arr[loc] == num {
			return arr[i] * arr[loc]
		}
	}
	return 0
}
