package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	east = iota
	south
	west
	north
)

// Runner runs the algorithm to get the answer
func Runner(arr []byte, num []int) int {
	var direction int
	var northDist, eastDist int
	for i := range arr {
		if arr[i] == 'F' {
			switch direction {
			case east:
				arr[i] = 'E'
			case south:
				arr[i] = 'S'
			case west:
				arr[i] = 'W'
			case north:
				arr[i] = 'N'
			}
		}
		switch arr[i] {
		case 'N':
			northDist += num[i]
		case 'S':
			northDist -= num[i]
		case 'E':
			eastDist += num[i]
		case 'W':
			eastDist -= num[i]
		case 'L':
			direction -= num[i] / 90
			if direction < 0 {
				direction += 4
			}
		case 'R':
			direction += num[i] / 90
			if direction > north {
				direction -= 4
			}
		}
		// fmt.Printf("east %d, north %d\n", eastDist, northDist)
	}
	return abs(northDist) + abs(eastDist)
}

func abs(in int) int {
	if in < 0 {
		return in * -1
	}
	return in
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) ([]byte, []int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []byte
	var num []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t[0])
		n, err := strconv.Atoi(t[1:])
		if err != nil {
			log.Fatal(err)
		}
		num = append(num, n)
	}
	return arr, num
}
