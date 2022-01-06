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
	wpEast := 10
	wpNorth := 1
	var distNorth, distEast int
	for i := range arr {
		switch arr[i] {
		case 'N':
			wpNorth += num[i]
		case 'S':
			wpNorth -= num[i]
		case 'E':
			wpEast += num[i]
		case 'W':
			wpEast -= num[i]
		case 'F':
			distEast += num[i] * wpEast
			distNorth += num[i] * wpNorth
		case 'L':
			for times := num[i] / 90; times > 0; times-- {
				wpEast, wpNorth = wpNorth*-1, wpEast
			}
		case 'R':
			for times := num[i] / 90; times > 0; times-- {
				wpEast, wpNorth = wpNorth, wpEast*-1
			}
		}
		// fmt.Printf("waypoint %d east, %d north; east %d, north %d\n", wpEast, wpNorth, distEast, distNorth)
	}
	return abs(distNorth) + abs(distEast)
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
