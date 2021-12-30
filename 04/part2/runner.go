package aoc

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Runner runs the algorithm to get the answer
func Runner(arr []map[string]string) int {
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
	var ans int
	for i := range arr {
		// fmt.Println(i)
		v, ok := arr[i]["byr"]
		if !ok {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		if n < 1920 || n > 2002 {
			continue
		}

		v, ok = arr[i]["iyr"]
		if !ok {
			continue
		}
		n, err = strconv.Atoi(v)
		if err != nil {
			continue
		}
		if n < 2010 || n > 2020 {
			continue
		}

		v, ok = arr[i]["eyr"]
		if !ok {
			continue
		}
		n, err = strconv.Atoi(v)
		if err != nil {
			continue
		}
		if n < 2020 || n > 2030 {
			continue
		}

		v, ok = arr[i]["hgt"]
		if !ok {
			continue
		}
		if strings.HasSuffix(v, "cm") {
			n, err = strconv.Atoi(v[:len(v)-2])
			if err != nil {
				continue
			}
			if n < 150 || n > 193 {
				continue
			}
		} else if strings.HasSuffix(v, "in") {
			n, err = strconv.Atoi(v[:len(v)-2])
			if err != nil {
				continue
			}
			if n < 59 || n > 76 {
				continue
			}
		} else {
			continue
		}

		v, ok = arr[i]["hcl"]
		if !ok {
			continue
		}
		validID := regexp.MustCompile(`^[#]{1}[a-f0-9]{6}$`)
		if !validID.MatchString(v) {
			continue
		}

		v, ok = arr[i]["ecl"]
		if !ok {
			continue
		}
		switch v {
		case "amb":
		case "blu":
		case "brn":
		case "gry":
		case "grn":
		case "hzl":
		case "oth":
		default:
			continue
		}

		v, ok = arr[i]["pid"]
		if !ok {
			continue
		}
		validID = regexp.MustCompile(`^[0-9]{9}$`)
		if !validID.MatchString(v) {
			continue
		}
		// fmt.Println("valid")
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
