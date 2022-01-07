package aoc

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(ts int, arr []string) int {
	bus := make([]int, 0, len(arr))
	for i := range arr {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			continue
		}
		bus = append(bus, num)
	}
	var id int
	wait := math.MaxInt32
	for i := range bus {
		rem := ts % bus[i]
		next := bus[i] - rem
		if next < wait {
			id = bus[i]
			wait = next
		}
	}
	return id * wait
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
