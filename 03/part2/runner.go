package aoc

import (
	"bufio"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(filename string) int {
	arr := parseFile(filename)
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }

	// fmt.Println(treeCounter(arr, 1, 1))
	// fmt.Println(treeCounter(arr, 3, 1))
	// fmt.Println(treeCounter(arr, 5, 1))
	// fmt.Println(treeCounter(arr, 7, 1))
	// fmt.Println(treeCounter(arr, 1, 2))
	ans := treeCounter(arr, 1, 1) * treeCounter(arr, 3, 1) * treeCounter(arr, 5, 1) * treeCounter(arr, 7, 1) * treeCounter(arr, 1, 2)
	return ans
}

func treeCounter(tree [][]bool, right, down int) int {
	var ans int
	for i := 1; i < len(tree); i++ {
		x := right * i
		x = x % len(tree[0])
		y := i * down
		if y >= len(tree) {
			break
		}
		if tree[y][x] {
			ans++
		}
	}
	return ans
}

func parseFile(filename string) [][]bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr [][]bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		barr := make([]bool, len(t))
		for i := range t {
			if t[i] == '#' {
				barr[i] = true
			}
		}
		arr = append(arr, barr)
	}
	return arr
}
