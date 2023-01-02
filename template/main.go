package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// parseFile reads the file and converts it to a format to use
func parseFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	type line struct {
		first  int
		second int
		letter byte
		text   string
	}

	var arr []int
	for {
		var l line
		_, err := fmt.Fscanf(file, "%d-%d %c: %s", &l.first, &l.second, &l.letter, &l.text)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		arr = append(arr, l.first)
	}
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	t := scanner.Text()
	// }
	return arr
}

func part1(filename string) int {
	input := parseFile(filename)
	fmt.Println(input)
	var answer int
	return answer
}

func part2(filename string) int {
	var answer int
	return answer
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 0 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	// if answer := part2(testFile); answer != 0 {
	// 	log.Fatalln("Wrong answer, got", answer)
	// }

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	// fmt.Println(part2(actualFile))
}
