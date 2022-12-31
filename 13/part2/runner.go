package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(_ int, arr []string) int {
	bus := make([]int, len(arr))
	for i := range arr {
		if arr[i] == "x" {
			continue
		}
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			log.Fatal(err)
		}
		bus[i] = num
	}
	var pos int
	for i := range bus {
		// fmt.Println(i, " ", bus[i])
		if bus[i] > bus[pos] {
			pos = i
		}
	}
	for i := 1; ; i++ {
		var notfound bool
		start := bus[pos]*i - pos
		for j := range bus {
			if bus[j] == 0 {
				continue
			}
			if (start+j)%bus[j] != 0 {
				notfound = true
				break
			}
		}
		if !notfound {
			return start
		}
	}
	return 0
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) (int, []string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	t := scanner.Text()
	ts, err := strconv.Atoi(t)
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	t = scanner.Text()
	arr := strings.Split(t, ",")
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
	return ts, arr
}

// type line struct {
// 	first  int
// 	second int
// 	letter byte
// 	text   string
// }
