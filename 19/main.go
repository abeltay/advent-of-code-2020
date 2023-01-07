package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

type rule struct {
	set1 []int
	set2 []int
	val  string
}

func parseInput(filename string) (map[int]rule, []string) {
	input := parseFile(filename)

	rules := make(map[int]rule)
	var pointer int
	for ; input[pointer] != ""; pointer++ {
		s := strings.Split(input[pointer], " ")
		num, err := strconv.Atoi(s[0][:len(s[0])-1])
		if err != nil {
			log.Fatalln("num conversion", err)
		}
		var secondSet bool
		var r rule
		for i := 1; i < len(s); i++ {
			if s[i] == "|" {
				secondSet = true
				continue
			}
			if s[i][0] == '"' {
				r.val = string(s[i][1])
				break
			}
			num, err := strconv.Atoi(s[i])
			if err != nil {
				log.Fatalln(err)
			}
			if !secondSet {
				r.set1 = append(r.set1, num)
			} else {
				r.set2 = append(r.set2, num)
			}
		}
		rules[num] = r
	}
	pointer++
	var messages []string
	for ; pointer < len(input); pointer++ {
		messages = append(messages, input[pointer])
	}
	return rules, messages
}

func generateForSet(rules map[int]rule, generated map[int][]string, set []int) []string {
	if set == nil {
		return nil
	}
	var possibilities []string
	for _, v := range set {
		if generated[v] == nil {
			generatePossiblities(rules, generated, v)
		}
		cur := generated[v]
		if len(possibilities) == 0 {
			possibilities = cur
			continue
		}
		var newArr []string
		for _, s := range possibilities {
			for _, s1 := range cur {
				newArr = append(newArr, s+s1)
			}
		}
		possibilities = newArr
	}
	return possibilities
}

func generatePossiblities(rules map[int]rule, generated map[int][]string, rule int) {
	if generated[rule] != nil {
		return
	}
	if rules[rule].val != "" {
		generated[rule] = []string{rules[rule].val}
		return
	}
	sets := generateForSet(rules, generated, rules[rule].set1)
	sets = append(sets, generateForSet(rules, generated, rules[rule].set2)...)
	generated[rule] = sets
}

func part1(filename string) int {
	rules, messages := parseInput(filename)
	generated := make(map[int][]string)
	generatePossiblities(rules, generated, 0)
	values := generated[0]
	sort.Strings(values)
	var answer int
	for _, v := range messages {
		idx := sort.SearchStrings(values, v)
		if idx < len(values) && values[idx] == v {
			answer++
		}
	}
	return answer
}

func combine8and11(fortyTwo, thirtyOne map[string]bool, length int, msg string) bool {
	var segments []string
	for i := 0; i < len(msg); i += length {
		segments = append(segments, msg[i:i+length])
	}
	front := 0
	for ; front < len(segments); front++ {
		if !fortyTwo[segments[front]] {
			break
		}
	}
	back := 0
	for ; back < len(segments); back++ {
		if !thirtyOne[segments[len(segments)-1-back]] {
			break
		}
	}
	if back == 0 {
		return false
	}
	if front+back < len(segments) {
		return false
	}
	return front > back
}

func part2(filename string, length int) int {
	rules, messages := parseInput(filename)
	generated := make(map[int][]string)
	generatePossiblities(rules, generated, 42)
	generatePossiblities(rules, generated, 31)
	map42 := make(map[string]bool)
	for _, v := range generated[42] {
		map42[v] = true
	}
	map31 := make(map[string]bool)
	for _, v := range generated[31] {
		map31[v] = true
	}
	var answer int
	for _, v := range messages {
		if combine8and11(map42, map31, length, v) {
			answer++
		}
	}
	return answer
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 2 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	testFile2 := "input2_test.txt"
	if answer := part1(testFile2); answer != 3 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile2, 5); answer != 12 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	// Combinations are 8 characters each
	fmt.Println(part2(actualFile, 8))
}
