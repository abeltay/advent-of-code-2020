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
	sort.Ints(arr)
	var prev, diff1, diff3 int
	for i := 0; i < len(arr); i++ {
		switch arr[i] - prev {
		case 1:
			diff1++
		case 3:
			diff3++
		default:
			log.Fatal("Diff error")
		}
		prev = arr[i]
	}
	return diff1 * (diff3 + 1)
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
