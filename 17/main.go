package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type location struct {
	x int
	y int
	z int
	w int
}

func parseFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	var arr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t)
	}
	return arr
}

func parseInput(filename string) map[location]bool {
	input := parseFile(filename)
	arr := make(map[location]bool)
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				arr[location{x: x, y: y}] = true
			}
		}
	}
	return arr
}

func size(cubes map[location]bool) (location, location) {
	var min, max location
	for k, v := range cubes {
		if v {
			min = k
			max = k
		}
	}
	for v := range cubes {
		if v.x < min.x {
			min.x = v.x
		}
		if v.x > max.x {
			max.x = v.x
		}
		if v.y < min.y {
			min.y = v.y
		}
		if v.y > max.y {
			max.y = v.y
		}
		if v.z < min.z {
			min.z = v.z
		}
		if v.z > max.z {
			max.z = v.z
		}
		if v.w < min.w {
			min.w = v.w
		}
		if v.w > max.w {
			max.w = v.w
		}
	}
	min.x--
	max.x++
	min.y--
	max.y++
	min.z--
	max.z++
	min.w--
	max.w++
	return min, max
}

func isActiveNext(cubes map[location]bool, current location) bool {
	var activeNeighbours int
	for x := current.x - 1; x <= current.x+1; x++ {
		for y := current.y - 1; y <= current.y+1; y++ {
			for z := current.z - 1; z <= current.z+1; z++ {
				l := location{
					x: x,
					y: y,
					z: z,
				}
				if current == l {
					continue
				}
				if cubes[l] {
					activeNeighbours++
				}
			}
		}
	}
	if cubes[current] {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			return true
		}
		return false
	}
	if activeNeighbours == 3 {
		return true
	}
	return false
}

func simulate(cubes map[location]bool) map[location]bool {
	min, max := size(cubes)
	newActive := make(map[location]bool)
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				loc := location{
					x: x,
					y: y,
					z: z,
				}
				if isActiveNext(cubes, loc) {
					newActive[loc] = true
				}
			}
		}
	}
	return newActive
}

func part1(filename string) int {
	input := parseInput(filename)
	for i := 0; i < 6; i++ {
		input = simulate(input)
	}
	return len(input)
}

func isActiveNext2(cubes map[location]bool, current location) bool {
	var activeNeighbours int
	for x := current.x - 1; x <= current.x+1; x++ {
		for y := current.y - 1; y <= current.y+1; y++ {
			for z := current.z - 1; z <= current.z+1; z++ {
				for w := current.w - 1; w <= current.w+1; w++ {
					l := location{
						x: x,
						y: y,
						z: z,
						w: w,
					}
					if current == l {
						continue
					}
					if cubes[l] {
						activeNeighbours++
					}
				}
			}
		}
	}
	if cubes[current] {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			return true
		}
		return false
	}
	if activeNeighbours == 3 {
		return true
	}
	return false
}

func simulate2(cubes map[location]bool) map[location]bool {
	min, max := size(cubes)
	newActive := make(map[location]bool)
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for z := min.z; z <= max.z; z++ {
				for w := min.w; w <= max.w; w++ {
					loc := location{
						x: x,
						y: y,
						z: z,
						w: w,
					}
					if isActiveNext2(cubes, loc) {
						newActive[loc] = true
					}
				}
			}
		}
	}
	return newActive
}

func part2(filename string) int {
	input := parseInput(filename)
	for i := 0; i < 6; i++ {
		input = simulate2(input)
	}
	return len(input)
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 112 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != 848 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
