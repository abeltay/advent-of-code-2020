package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	var arr []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		for _, v := range t {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalln(err)
			}
			arr = append(arr, i)
		}
	}
	return arr
}

func createRing(input []int) *ring.Ring {
	r := ring.New(len(input))
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = input[i]
		r = r.Next()
	}
	return r
}

func nextLower(value int) int {
	value--
	if value <= 0 {
		return value + 9
	}
	return value
}

func move(r *ring.Ring) *ring.Ring {
	next := r.Next()
	dest := nextLower(r.Value.(int))
	for i := 0; i < 3; i++ {
		if next.Value.(int) == dest {
			dest = nextLower(dest)
			next = r
			i = -1
		}
		next = next.Next()
	}
	three := r.Unlink(3)
	next = r.Next()
	for next.Value.(int) != dest {
		next = next.Next()
	}
	next.Link(three)
	return r.Next()
}

func answerToString(r *ring.Ring) string {
	for r.Value.(int) != 1 {
		r = r.Next()
	}
	var ans string
	n := r.Len() - 1
	for i := 0; i < n; i++ {
		r = r.Next()
		ans += strconv.Itoa(r.Value.(int))
	}
	return ans
}

func part1(filename string, moves int) string {
	input := parseInput(filename)
	r := createRing(input)
	for i := 0; i < moves; i++ {
		r = move(r)
	}
	return answerToString(r)
}

const cups = 1000000

func createRing2(input []int) (*ring.Ring, []*ring.Ring) {
	r := ring.New(cups)
	vec := make([]*ring.Ring, cups+1) // Not starting from 0
	for i := 0; i < len(input); i++ {
		r.Value = input[i]
		vec[input[i]] = r
		r = r.Next()
	}
	n := r.Len()
	for i := len(input) + 1; i <= n; i++ {
		r.Value = i
		vec[i] = r
		r = r.Next()
	}
	return r, vec
}

func nextLower2(value int) int {
	value--
	if value <= 0 {
		return value + cups
	}
	return value
}

func move2(r *ring.Ring, vec []*ring.Ring) *ring.Ring {
	next := r.Next()
	dest := nextLower2(r.Value.(int))
	for i := 0; i < 3; i++ {
		if next.Value.(int) == dest {
			dest = nextLower2(dest)
			next = r
			i = -1
		}
		next = next.Next()
	}
	three := r.Unlink(3)
	next = vec[dest]
	next.Link(three)
	return r.Next()
}

func part2(filename string) int {
	input := parseInput(filename)
	r, vec := createRing2(input)
	for i := 0; i < 10000000; i++ {
		r = move2(r, vec)
	}
	one := vec[1]
	return one.Next().Value.(int) * one.Next().Next().Value.(int)
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile, 10); answer != "92658374" {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part1(testFile, 100); answer != "67384529" {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != 149245887792 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile, 100))
	fmt.Println(part2(actualFile))
}
