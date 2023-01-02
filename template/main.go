package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func parseInput(filename string) []int {
	input := parseFile(filename)

	type line struct {
		first  int
		second int
		letter byte
		text   string
	}

	var arr []int
	for _, v := range input {
		var l line
		_, err := fmt.Sscanf(v, "%d-%d %c: %s", &l.first, &l.second, &l.letter, &l.text)
		if err != nil {
			log.Fatalln(err)
		}
		arr = append(arr, l.first)
	}
	return arr
}

func part1(filename string) int {
	input := parseInput(filename)
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
