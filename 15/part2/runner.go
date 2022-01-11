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
	seen := make(map[int][2]int)
	var pos [2]int
	for i := range arr {
		pos = [2]int{-1, i + 1}
		seen[arr[i]] = pos
	}
	last := arr[len(arr)-1]
	for i := len(seen) + 1; i <= 30000000; i++ {
		if pos[0] == -1 {
			last = 0
		} else {
			last = pos[1] - pos[0]
		}
		var ok bool
		pos, ok = seen[last]
		if ok {
			pos = [2]int{pos[1], i}
		} else {
			pos = [2]int{-1, i}
		}
		seen[last] = pos
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
