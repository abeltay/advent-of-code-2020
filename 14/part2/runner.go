package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []string) int {
	var mask string
	mem := make(map[int]int)
	for i := range arr {
		s := strings.Split(arr[i], " = ")
		if s[0] == "mask" {
			mask = s[1]
			continue
		}
		var loc int
		_, err := fmt.Sscanf(s[0], "mem[%d]", &loc)
		if err != nil {
			log.Fatal(err)
		}
		num := strconv.FormatInt(int64(loc), 2)
		newNum := []byte(mask)
		for j := range newNum {
			switch newNum[len(newNum)-1-j] {
			case '0':
				if j < len(num) {
					newNum[len(newNum)-1-j] = num[len(num)-1-j]
				} else {
					newNum[len(newNum)-1-j] = '0'
				}
			case '1':
				newNum[len(newNum)-1-j] = '1'
			case 'X':
				newNum[len(newNum)-1-j] = 'X'
			}
		}

		s1, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		arrs := resolveX(string(newNum))
		for i := range arrs {
			n1, err := strconv.ParseInt(arrs[i], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			mem[int(n1)] = int(s1)
			// fmt.Println(n1, s1)
		}
	}
	var ans int
	for i := range mem {
		ans += mem[i]
	}
	return ans
}

func resolveX(mask string) []string {
	variations := []string{
		mask,
	}
	for i := range mask {
		if mask[i] != 'X' {
			continue
		}
		newvar := make([]string, 0, len(variations)*2)
		for j := range variations {
			n := []byte(variations[j])
			n[i] = '0'
			newvar = append(newvar, string(n))
		}
		for j := range variations {
			n := []byte(variations[j])
			n[i] = '1'
			newvar = append(newvar, string(n))
		}
		variations = newvar
	}
	return variations
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t)
	}
	return arr
}
