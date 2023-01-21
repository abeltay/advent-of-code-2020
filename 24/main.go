package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func parseInput(filename string) [][]string {
	input := parseFile(filename)

	var arr [][]string
	for _, v := range input {
		var direction []string
		for i := 0; i < len(v); i++ {
			switch v[i] {
			case 'e':
				fallthrough
			case 'w':
				direction = append(direction, string(v[i]))
			default:
				direction = append(direction, string(v[i:i+2]))
				i++
			}
		}
		arr = append(arr, direction)
	}
	return arr
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func unkey(key string) (int, int) {
	s := strings.Split(key, ",")
	x, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatalln(err)
	}
	y, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatalln(err)
	}
	return x, y
}

func move(x, y int, direction string) (int, int) {
	switch direction {
	case "e":
		x++
	case "w":
		x--
	case "se":
		y--
		if y%2 == 0 {
			x++
		}
	case "sw":
		y--
		if y%2 != 0 {
			x--
		}
	case "nw":
		y++
		if y%2 != 0 {
			x--
		}
	case "ne":
		y++
		if y%2 == 0 {
			x++
		}
	}
	return x, y
}

func findCoordinate(direction []string) (int, int) {
	var x, y int
	for _, v := range direction {
		x, y = move(x, y, v)
	}
	return x, y
}

func flippedTiles(directions [][]string) map[string]bool {
	blackTile := make(map[string]bool)
	for _, v := range directions {
		x, y := findCoordinate(v)
		k := key(x, y)
		if blackTile[k] {
			blackTile[k] = false
		} else {
			blackTile[k] = true
		}
	}
	return blackTile
}

func countBlackTiles(blackTile map[string]bool) int {
	var answer int
	for _, v := range blackTile {
		if v {
			answer++
		}
	}
	return answer
}

func part1(filename string) int {
	input := parseInput(filename)
	blackTile := flippedTiles(input)
	return countBlackTiles(blackTile)
}

func adjacentBlacks(x, y int, blackTile map[string]bool) int {
	var count int
	if blackTile[key(move(x, y, "e"))] {
		count++
	}
	if blackTile[key(move(x, y, "se"))] {
		count++
	}
	if blackTile[key(move(x, y, "sw"))] {
		count++
	}
	if blackTile[key(move(x, y, "w"))] {
		count++
	}
	if blackTile[key(move(x, y, "nw"))] {
		count++
	}
	if blackTile[key(move(x, y, "ne"))] {
		count++
	}
	return count
}

func autoFlip(blackTile map[string]bool) map[string]bool {
	var minX, maxX, minY, maxY int
	for k, v := range blackTile {
		if v {
			x, y := unkey(k)
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	newBlackTile := make(map[string]bool)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			b := adjacentBlacks(x, y, blackTile)
			k := key(x, y)
			if blackTile[key(x, y)] {
				if b > 0 && b <= 2 {
					newBlackTile[k] = true
				}
			} else {
				if b == 2 {
					newBlackTile[k] = true
				}
			}
		}
	}
	return newBlackTile
}

func part2(filename string) int {
	input := parseInput(filename)
	blackTile := flippedTiles(input)
	for i := 1; i <= 100; i++ {
		blackTile = autoFlip(blackTile)
	}
	return countBlackTiles(blackTile)
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 10 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != 2208 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
