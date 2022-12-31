package aoc

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(arr []int) int {
	cache := make(map[int]int)
	arr = append(arr, 0)
	sort.Ints(arr)
	return findWays(arr, cache, 0)
}

func findWays(arr []int, cache map[int]int, cur int) int {
	if cur == len(arr)-1 {
		return 1
	}
	var total int
	for i := cur + 1; i < len(arr) && i <= cur+3; i++ {
		if arr[cur]+3 < arr[i] {
			break
		}
		num, ok := cache[i]
		if ok {
			total += num
			continue
		}
		total += findWays(arr, cache, i)
	}
	cache[cur] = total
	return total
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		num, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, num)
	}
	return arr
}
