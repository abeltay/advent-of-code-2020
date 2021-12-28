package aoc

import (
	"log"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(data []string) int {
	var ans int
	for i := range data {
		var min, max int
		var err error
		var letter byte
		s := strings.Split(data[i], " ")
		snum := strings.Split(s[0], "-")
		min, err = strconv.Atoi(snum[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err = strconv.Atoi(snum[1])
		if err != nil {
			log.Fatal(err)
		}
		letter = s[1][0]
		var count int
		for j := range s[2] {
			if s[2][j] == letter {
				count++
			}
		}
		if min <= count && count <= max {
			ans++
		}
	}
	return ans
}
