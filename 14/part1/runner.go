package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []string) int {
	var mask string
	mem := make(map[int]int)
	for i := range arr {
		s := strings.Split(arr[i], " = ")
		if s[0] == "mask" {
			mask = s[1]
			continue
		}
		s1, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		num := strconv.FormatInt(int64(s1), 2)
		newNum := []byte(mask)
		// for j := len(newNum) - 1; j >= 0; j-- {
		for j := range newNum {
			if newNum[len(newNum)-1-j] == 'X' {
				if j < len(num) {
					newNum[len(newNum)-1-j] = num[len(num)-1-j]
				} else {
					newNum[len(newNum)-1-j] = '0'
				}
			}
		}
		n1, err := strconv.ParseInt(string(newNum), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		var loc int
		_, err = fmt.Sscanf(s[0], "mem[%d]", &loc)
		if err != nil {
			log.Fatal(err)
		}
		mem[loc] = int(n1)
	}
	var ans int
	for i := range mem {
		ans += mem[i]
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
