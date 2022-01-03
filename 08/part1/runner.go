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
	var ans, pointer int
	for pointer < len(arr) {
		if seen[pointer] {
			break
		}
		seen[pointer] = true
		switch arr[pointer].op {
		case "acc":
			ans += arr[pointer].val
			fallthrough
		case "nop":
			pointer++
		case "jmp":
			pointer += arr[pointer].val
		}
	}
	return ans
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
