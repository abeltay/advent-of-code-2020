package aoc

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []map[string]string) int {
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
	var ans int
	for i := range arr {
		if _, ok := arr[i]["byr"]; !ok {
			continue
		}
		if _, ok := arr[i]["iyr"]; !ok {
			continue
		}
		if _, ok := arr[i]["eyr"]; !ok {
			continue
		}
		if _, ok := arr[i]["hgt"]; !ok {
			continue
		}
		if _, ok := arr[i]["hcl"]; !ok {
			continue
		}
		if _, ok := arr[i]["ecl"]; !ok {
			continue
		}
		if _, ok := arr[i]["pid"]; !ok {
			continue
		}
		ans++
	}
	return ans
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []map[string]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []map[string]string
	nmap := make(map[string]string)
	arr = append(arr, nmap)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			nmap = make(map[string]string)
			arr = append(arr, nmap)
			continue
		}
		s := strings.Split(t, " ")
		for _, v := range s {
			ns := strings.Split(v, ":")
			nmap[ns[0]] = ns[1]
		}
	}
	return arr
}
