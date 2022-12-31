package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []int) int {
	prev := make(map[int]int)
	seen := make(map[int]int)
	var last int
	for i := range arr {
		seen[arr[i]] = i + 1
		last = arr[i]
	}
	for i := len(seen) + 1; i <= 2020; i++ {
		pos2, ok := prev[last]
		if !ok {
			last = 0
			prev[last], seen[last] = seen[last], i
		} else {
			pos1 := seen[last]
			last = pos1 - pos2
			if _, ok = seen[last]; ok {
				prev[last] = seen[last]
			}
			seen[last] = i
		}
		// fmt.Printf("turn: %d, last: %d\n", i, last)
	}
	return last
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
		s := strings.Split(t, ",")
		for i := range s {
			num, err := strconv.Atoi(s[i])
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, num)
		}
	}
	return arr
}
