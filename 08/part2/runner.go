package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []line) int {
	seen := make([]bool, len(arr))
	var acc, p int
	for p < len(arr) {
		seen[p] = true
		switch arr[p].op {
		case "acc":
			acc += arr[p].val
			p++
		case "nop":
			t := checkTermination(arr, seen, p+arr[p].val, acc)
			if t != 0 {
				return t
			}
			p++
		case "jmp":
			t := checkTermination(arr, seen, p+1, acc)
			if t != 0 {
				return t
			}
			p += arr[p].val
		}
	}
	return acc
}

func checkTermination(arr []line, seen []bool, p, acc int) int {
	for p < len(arr) {
		if seen[p] {
			return 0
		}
		seen[p] = true
		switch arr[p].op {
		case "acc":
			acc += arr[p].val
			p++
		case "nop":
			p++
		case "jmp":
			p += arr[p].val
		}
	}
	return acc
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []line {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		s := strings.Split(t, " ")
		num, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		l := line{
			op:  s[0],
			val: num,
		}
		arr = append(arr, l)
	}
	return arr
}

type line struct {
	op  string
	val int
}
