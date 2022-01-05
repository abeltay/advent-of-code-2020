package aoc

import (
	"bufio"
	"log"
	"os"
)

// Runner runs the algorithm to get the answer
func Runner(arr [][]byte) int {
	// for i := range arr {
	// 	for j := range arr[i] {
	// 		fmt.Print(string(arr[i][j]))
	// 	}
	// 	fmt.Println()
	// }
	moved := true
	for moved {
		visibleSeats := updateSight(arr)
		arr, moved = occupy(arr, visibleSeats)
		// fmt.Println("----------")
		// for i := range arr {
		// 	for j := range arr[i] {
		// 		fmt.Print(string(arr[i][j]))
		// 	}
		// 	fmt.Println()
		// }
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

func occupy(arr [][]byte, visibleSeats [][]int) ([][]byte, bool) {
	var moved bool
	var newArr [][]byte
	for i := range arr {
		n := make([]byte, len(arr[i]))
		for j := range arr[i] {
			if arr[i][j] == 'L' && visibleSeats[i][j] == 0 {
				n[j] = '#'
				moved = true
			} else if arr[i][j] == '#' && visibleSeats[i][j] >= 5 {
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

func updateSight(arr [][]byte) [][]int {
	var visibleSeats [][]int
	for i := range arr {
		n := make([]int, len(arr[i]))
		visibleSeats = append(visibleSeats, n)
	}
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == '#' {
				for k := 1; i-k >= 0 && j-k >= 0; k++ {
					if arr[i-k][j-k] != '.' {
						visibleSeats[i-k][j-k]++
						break
					}
				}
				for k := 1; i-k >= 0; k++ {
					if arr[i-k][j] != '.' {
						visibleSeats[i-k][j]++
						break
					}
				}
				for k := 1; i-k >= 0 && j+k < len(arr[i]); k++ {
					if arr[i-k][j+k] != '.' {
						visibleSeats[i-k][j+k]++
						break
					}
				}
				for k := 1; j-k >= 0; k++ {
					if arr[i][j-k] != '.' {
						visibleSeats[i][j-k]++
						break
					}
				}
				for k := 1; j+k < len(arr[i]); k++ {
					if arr[i][j+k] != '.' {
						visibleSeats[i][j+k]++
						break
					}
				}
				for k := 1; i+k < len(arr) && j-k >= 0; k++ {
					if arr[i+k][j-k] != '.' {
						visibleSeats[i+k][j-k]++
						break
					}
				}
				for k := 1; i+k < len(arr); k++ {
					if arr[i+k][j] != '.' {
						visibleSeats[i+k][j]++
						break
					}
				}
				for k := 1; i+k < len(arr) && j+k < len(arr[i]); k++ {
					if arr[i+k][j+k] != '.' {
						visibleSeats[i+k][j+k]++
						break
					}
				}
			}
		}
	}
	return visibleSeats
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
