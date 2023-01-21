package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var arr []int
	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		arr = append(arr, i)
	}
	return arr
}

func steps(subject, value int) int {
	return (subject * value) % 20201227
}

func findLoopSize(pkey int) int {
	value := 1
	var loop int
	for value != pkey {
		value = steps(7, value)
		loop++
	}
	return loop
}

func part1(filename string) int {
	input := parseInput(filename)
	loop := findLoopSize(input[0])
	value := 1
	for i := 0; i < loop; i++ {
		value = steps(input[1], value)
	}
	return value
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 14897079 {
		log.Fatalln("Wrong answer, got:", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
}
