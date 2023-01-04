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

func operate(first, second, operation int) int {
	if operation == 0 {
		return first + second
	} else {
		return first * second
	}
}

func evaluate(line string) (int, int) {
	pointer := 1
	var value int
	var operation int
	for pointer < len(line) {
		switch line[pointer] {
		case '+':
			operation = 0
		case '*':
			operation = 1
		case '(':
			second, newPointer := evaluate(line[pointer:])
			pointer += newPointer
			value = operate(value, second, operation)
		case ')':
			return value, pointer
		case ' ':
		default: // number
			s := line[pointer]
			pointer++
			for line[pointer] >= '0' && line[pointer] <= '9' {
				s += line[pointer]
				pointer++
			}
			pointer--
			second, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatalln(err)
			}
			value = operate(value, second, operation)
		}
		pointer++
	}
	return value, pointer
}

func part1(filename string) int {
	input := parseFile(filename)
	var answer int
	for _, v := range input {
		ans, _ := evaluate("(" + v + ")")
		answer += ans
	}
	return answer
}

func parseFile2(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	var arr [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, strings.Split(strings.ReplaceAll(t, " ", ""), ""))
	}
	return arr
}

func calculateNoBrackets(input []string) string {
	for i := 0; i < len(input); i++ {
		if input[i] == "+" {
			val1, err := strconv.Atoi(input[i-1])
			if err != nil {
				log.Fatalln("1", err)
			}
			val2, err := strconv.Atoi(input[i+1])
			if err != nil {
				log.Fatalln("2", err)
			}
			input = append(input[:i], input[i+2:]...)
			input[i-1] = strconv.Itoa(val1 + val2)
			i--
		}
	}
	for i := 0; i < len(input); i++ {
		if input[i] == "*" {
			val1, err := strconv.Atoi(input[i-1])
			if err != nil {
				log.Fatalln("3", err)
			}
			val2, err := strconv.Atoi(input[i+1])
			if err != nil {
				log.Fatalln("4", err)
			}
			input = append(input[:i], input[i+2:]...)
			input[i-1] = strconv.Itoa(val1 * val2)
			i--
		}
	}
	return input[0]
}

func evaluate2(input []string) int {
	var bracketStack []int
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case "(":
			bracketStack = append(bracketStack, i)
		case ")":
			start := bracketStack[len(bracketStack)-1]
			s := calculateNoBrackets(input[start+1 : i])
			input = append(input[:start], input[i:]...)
			input[start] = s
			i = start
			bracketStack = bracketStack[:len(bracketStack)-1]
		}
	}
	ansStr := calculateNoBrackets(input)
	ans, err := strconv.Atoi(ansStr)
	if err != nil {
		log.Println(ans, err)
	}
	return ans
}

func part2(filename string) int {
	input := parseFile2(filename)
	var answer int
	for _, v := range input {
		ans := evaluate2(v)
		answer += ans
	}
	return answer
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 26386 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != 693942 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
