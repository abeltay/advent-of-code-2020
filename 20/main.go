package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type tile struct {
	id    int
	style [][]int
}

func parseInput(filename string) []tile {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	var tiles []tile
	var t1 tile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		switch {
		case t == "":
			tiles = append(tiles, t1)
			t1 = tile{}
		case t[:4] == "Tile":
			_, err := fmt.Sscanf(t, "Tile %d:", &t1.id)
			if err != nil {
				log.Fatalln(err)
			}
		default:
			s := strings.Split(t, "")
			var arr []int
			for _, v := range s {
				if v == "#" {
					arr = append(arr, 1)
				} else {
					arr = append(arr, 0)
				}
			}
			t1.style = append(t1.style, arr)
		}
	}
	return tiles
}

func flipTileVertical(t [][]int) [][]int {
	var flipped [][]int
	for y := range t {
		row := make([]int, len(t[y]))
		for x := range t[y] {
			row[x] = t[len(t)-1-y][x]
		}
		flipped = append(flipped, row)
	}
	return flipped
}

func rotateTileClockwise(t [][]int) [][]int {
	var rotated [][]int
	for y := range t {
		row := make([]int, len(t[y]))
		for x := range t {
			row[x] = t[len(t[0])-1-x][y]
		}
		rotated = append(rotated, row)
	}
	return rotated
}

func addAllRotations(t [][]int) [][][]int {
	tiles := make([][][]int, 0, 4)
	tiles = append(tiles, t)
	rotated := t
	for i := 0; i < 3; i++ {
		rotated = rotateTileClockwise(rotated)
		tiles = append(tiles, rotated)
	}
	return tiles
}

func generatePermutations(t [][]int) [][][]int {
	tiles := make([][][]int, 0, 8)
	tiles = append(tiles, addAllRotations(t)...)
	tiles = append(tiles, addAllRotations(flipTileVertical(t))...)
	return tiles
}

func isSideMatched(left, right tile) bool {
	for y := range left.style {
		if left.style[y][len(left.style[y])-1] != right.style[y][0] {
			return false
		}
	}
	return true
}

func matchLeft(fixed, match tile) *tile {
	matches := generatePermutations(match.style)
	for _, v := range matches {
		t := tile{
			id:    match.id,
			style: v,
		}
		if isSideMatched(t, fixed) {
			return &t
		}
	}
	return nil
}

func matchRight(fixed, match tile) *tile {
	matches := generatePermutations(match.style)
	for _, v := range matches {
		t := tile{
			id:    match.id,
			style: v,
		}
		if isSideMatched(fixed, t) {
			return &t
		}
	}
	return nil
}

func isTopMatched(top, bottom tile) bool {
	for x := range top.style {
		if top.style[len(top.style)-1][x] != bottom.style[0][x] {
			return false
		}
	}
	return true
}

func matchTop(fixed, match tile) *tile {
	matches := generatePermutations(match.style)
	for _, v := range matches {
		t := tile{
			id:    match.id,
			style: v,
		}
		if isTopMatched(t, fixed) {
			return &t
		}
	}
	return nil
}

func matchBottom(fixed, match tile) *tile {
	matches := generatePermutations(match.style)
	for _, v := range matches {
		t := tile{
			id:    match.id,
			style: v,
		}
		if isTopMatched(fixed, t) {
			return &t
		}
	}
	return nil
}

func buildLine(tiles, line []tile) ([]tile, []tile) {
	for i := 0; i < len(tiles); i++ {
		t := matchRight(line[len(line)-1], tiles[i])
		if t != nil {
			line = append(line, *t)
			tiles = append(tiles[:i], tiles[i+1:]...)
			i = -1
			continue
		}
	}
	return tiles, line
}

func buildImage(tiles []tile) [][]tile {
	// size := len(tiles)
	zero := tile{
		id:    tiles[0].id,
		style: flipTileVertical(tiles[0].style),
	}
	line := []tile{zero}
	tiles = tiles[1:]
	// Build a line
	for i := 0; i < len(tiles); i++ {
		t := matchLeft(line[0], tiles[i])
		if t != nil {
			line = append([]tile{*t}, line...)
			tiles = append(tiles[:i], tiles[i+1:]...)
			i = -1
			continue
		}
	}
	tiles, line = buildLine(tiles, line)
	var image [][]tile
	image = append(image, line)
	for i := 0; i < len(tiles); i++ {
		t := matchBottom(image[len(image)-1][0], tiles[i])
		if t != nil {
			// printTile(*t)
			tiles = append(tiles[:i], tiles[i+1:]...)
			tiles, line = buildLine(tiles, []tile{*t})
			image = append(image, line)
			i = -1
			continue
		}
		t = matchTop(image[0][0], tiles[i])
		if t != nil {
			// printTile(*t)
			tiles = append(tiles[:i], tiles[i+1:]...)
			tiles, line = buildLine(tiles, []tile{*t})
			image = append([][]tile{line}, image...)
			i = -1
			continue
		}
	}
	return image
}

func part1(filename string) int {
	input := parseInput(filename)
	image := buildImage(input)
	answer := image[0][0].id * image[0][len(image[0])-1].id * image[len(image)-1][0].id * image[len(image)-1][len(image[0])-1].id
	return answer
}

func stripBorders(image [][]int) [][]int {
	var actual [][]int
	for y := 1; y < len(image)-1; y++ {
		actual = append(actual, image[y][1:len(image[y])-1])
	}
	return actual
}

func buildRow(image []tile) [][]int {
	actual := stripBorders(image[0].style)
	for i := 1; i < len(image); i++ {
		n := stripBorders(image[i].style)
		for j := range actual {
			actual[j] = append(actual[j], n[j]...)
		}
	}
	return actual
}

func buildActualImage(image [][]tile) [][]int {
	var actual [][]int
	for y := range image {
		actual = append(actual, buildRow(image[y])...)
	}
	return actual
}

var seaMonster = [][]int{
	[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	[]int{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	[]int{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0},
}

func matchSeaMonster(img [][]int, coordX, coordY int) bool {
	for y := 0; y < len(seaMonster); y++ {
		for x := 0; x < len(seaMonster[y]); x++ {
			if seaMonster[y][x] == 1 && img[coordY+y][coordX+x] != 1 {
				return false
			}
		}
	}
	return true
}

func countSeaMonster(img [][]int) int {
	var count int
	for y := 0; y < len(img)-len(seaMonster); y++ {
		for x := 0; x < len(img[y])-len(seaMonster[0]); x++ {
			if matchSeaMonster(img, x, y) {
				count++
			}
		}
	}
	return count
}

func part2(filename string) int {
	input := parseInput(filename)
	image := buildImage(input)
	// printTile(image[len(image)-1][0])
	actual := buildActualImage(image)
	var answer int
	for y := range actual {
		for x := range actual[y] {
			if actual[y][x] == 1 {
				answer++
			}
		}
	}
	permutations := generatePermutations(actual)
	for _, v := range permutations {
		m := countSeaMonster(v)
		if m > 0 {
			return answer - (m * 15)
		}
	}
	return 0
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 20899048083289 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != 273 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
