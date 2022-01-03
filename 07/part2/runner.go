package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []line) int {
	// for i := range arr {
	// 	fmt.Println(arr[i])
	// }
	return bagDiver(arr, "shiny gold", 1) - 1
}

func bagDiver(arr []line, name string, num int) int {
	ans := num
	for i := range arr {
		if arr[i].bag == name {
			if len(arr[i].inner) == 0 {
				return num
			}
			for j := range arr[i].inner {
				ans += num * bagDiver(arr, arr[i].inner[j], arr[i].size[j])
			}
			break
		}
	}
	// fmt.Println(name, num, ans)
	return ans
}

type line struct {
	bag   string
	inner []string
	size  []int
}

// ParseFile reads the file and converts it to a format for runner to use
func ParseFile(filename string) []line {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q", filename, err)
	}
	defer f.Close()

	var arr []line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var l line
		t := scanner.Text()
		s := strings.Split(t, " contain ")
		s1 := strings.Split(s[0], " ")
		l.bag = s1[0] + " " + s1[1]
		s1 = strings.Split(s[1], ", ")
		for i := range s1 {
			if s1[i] == "no other bags." {
				continue
			}
			s2 := strings.Split(s1[i], " ")
			num, err := strconv.Atoi(s2[0])
			if err != nil {
				log.Fatal("conversion ", err)
			}
			l.size = append(l.size, num)
			l.inner = append(l.inner, s2[1]+" "+s2[2])
		}
		arr = append(arr, l)
	}
	return arr
}
