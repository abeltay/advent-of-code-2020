package aoc

import (
	"bufio"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(arr [][]string) int {
	var ans int
	for i := range arr {
		m := make(map[string]int)
		for j := range arr[i] {
			for k := range arr[i][j] {
				m[string(arr[i][j][k])]++
			}
		}
		for j := range m {
			if m[j] == len(arr[i]) {
				ans++
			}
		}
	}
	return ans
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr [][]string
	var cur []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			arr = append(arr, cur)
			cur = []string{}
			continue
		}
		cur = append(cur, t)
	}
	arr = append(arr, cur)
	return arr
}
