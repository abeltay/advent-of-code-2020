package aoc

import (
	"bufio"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(arr [][]byte) int {
	moved := true
	for moved {
		arr, moved = occupy(arr)
	}
	var ans int
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == '#' {
				ans++
			}
		}
	}
	return ans
}

func occupy(arr [][]byte) ([][]byte, bool) {
	var moved bool
	var newArr [][]byte
	for i := range arr {
		n := make([]byte, len(arr[i]))
		for j := range arr[i] {
			c := countAdj(arr, i, j)
			if arr[i][j] == 'L' && c == 0 {
				n[j] = '#'
				moved = true
			} else if arr[i][j] == '#' && c >= 4 {
				n[j] = 'L'
				moved = true
			} else {
				n[j] = arr[i][j]
			}
		}
		newArr = append(newArr, n)
	}
	return newArr, moved
}

func countAdj(arr [][]byte, i, j int) int {
	var count int
	for i1 := i - 1; i1 <= i+1; i1++ {
		if i1 < 0 || i1 >= len(arr) {
			continue
		}
		for j1 := j - 1; j1 <= j+1; j1++ {
			if j1 < 0 || j1 >= len(arr[0]) {
				continue
			}
			if i1 == i && j1 == j {
				continue
			}
			if arr[i1][j1] == '#' {
				count++
			}
		}
	}
	return count
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) [][]byte {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, []byte(t))
	}
	// var arr []line
	// for {
	// 	var l line
	// 	_, err := fmt.Fscanf(f, "%d-%d %c: %s\n", &l.first, &l.second, &l.letter, &l.text)
	// 	// _, err := fmt.Fscanln(f, &num)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	arr = append(arr, num)
	// }
	return arr
}

// type line struct {
// 	first  int
// 	second int
// 	letter byte
// 	text   string
// }
