package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	from1 int
	to1   int
	from2 int
	to2   int
}

// parseFile reads the file and converts it to a format to use
func parseFile(filename string) ([]rule, []int, [][]int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open %q: %q\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	var allRules []rule
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		s := strings.Split(t, ": ")
		var r rule
		_, err := fmt.Sscanf(s[1], "%d-%d or %d-%d", &r.from1, &r.to1, &r.from2, &r.to2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		allRules = append(allRules, r)
	}

	var myTicket []int
	scanner.Scan()
	scanner.Scan()
	ticket := scanner.Text()
	s := strings.Split(ticket, ",")
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		myTicket = append(myTicket, i)
	}

	var nearbyTickets [][]int
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		ticket := scanner.Text()
		s := strings.Split(ticket, ",")
		var current []int
		for _, v := range s {
			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			current = append(current, i)
		}
		nearbyTickets = append(nearbyTickets, current)
	}
	return allRules, myTicket, nearbyTickets
}

func valid(rule rule, field int) bool {
	if (field >= rule.from1 && field <= rule.to1) || (field >= rule.from2 && field <= rule.to2) {
		return true
	}
	return false
}

func invalidField(rules []rule, ticket []int) int {
	for _, field := range ticket {
		var matched bool
		for _, rule := range rules {
			if valid(rule, field) {
				matched = true
				break
			}
		}
		if !matched {
			return field
		}
	}
	return 0
}

func part1(filename string) int {
	var answer int
	rules, _, nearbyTickets := parseFile(filename)
	for _, ticket := range nearbyTickets {
		answer += invalidField(rules, ticket)
	}
	return answer
}

func filterValidTickets(rules []rule, nearbyTickets [][]int) [][]int {
	var validTickets [][]int
	for _, ticket := range nearbyTickets {
		if invalidField(rules, ticket) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func findValidField(rule rule, nearbyTickets [][]int) []int {
	var validFields []int
	for column := range nearbyTickets[0] {
		var foundInvalid bool
		for _, ticket := range nearbyTickets {
			if !valid(rule, ticket[column]) {
				foundInvalid = true
				break
			}
		}
		if !foundInvalid {
			validFields = append(validFields, column)
		}
	}
	return validFields
}

func top6found(fieldmap []int) bool {
	for i := 0; i < 6; i++ {
		if fieldmap[i] == -1 {
			return false
		}
	}
	return true
}

func simplify(fieldMap []int, validFields [][]int) {
	var field int
	for i := range validFields {
		if len(validFields[i]) == 1 {
			field = validFields[i][0]
			fieldMap[i] = field
		}
	}
	for i := range validFields {
		for j := range validFields[i] {
			if validFields[i][j] == field {
				validFields[i][j] = validFields[i][len(validFields[i])-1]
				validFields[i] = validFields[i][:len(validFields[i])-1]
				break
			}
		}
	}
}

func part2(filename string) int {
	rules, myTicket, nearbyTickets := parseFile(filename)
	validTickets := filterValidTickets(rules, nearbyTickets)
	validTickets = append(validTickets, myTicket)
	var validFields [][]int
	for i := range rules {
		value := findValidField(rules[i], validTickets)
		validFields = append(validFields, value)
	}
	fieldMap := make([]int, len(rules))
	for i := range fieldMap {
		fieldMap[i] = -1
	}
	for !top6found(fieldMap) {
		simplify(fieldMap, validFields)
	}
	answer := 1
	for i := 0; i < 6; i++ {
		answer *= myTicket[fieldMap[i]]
	}
	return answer
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 71 {
		fmt.Println("Wrong answer, got:", answer)
		return
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
