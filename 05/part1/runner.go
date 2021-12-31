package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Runner runs the algorithm to get the answer
func Runner(arr []string) int {
	var ans int
	for i := range arr {
		var seat string
		for j := range arr[i] {
			if arr[i][j] == 'F' || arr[i][j] == 'L' {
				seat += "0"
			} else if arr[i][j] == 'B' || arr[i][j] == 'R' {
				seat += "1"
			}
		}
		num, err := strconv.ParseInt(seat, 2, 64)
		if err != nil {
			log.Println(err)
		}
		if int(num) > ans {
			ans = int(num)
		}
	}
	return ans
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t)
	}
	return arr
}
