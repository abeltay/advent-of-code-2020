package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type food struct {
	ingredient []string
	allergen   []string
}

func parseInput(filename string) []food {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	var arr []food
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := scanner.Text()
		idx := strings.Index(v, " (contain")
		in := food{
			ingredient: strings.Split(v[:idx], " "),
			allergen:   strings.Split(v[idx+11:len(v)-1], ", "),
		}
		arr = append(arr, in)
	}
	return arr
}

func intersection(s1, s2 []string) []string {
	var out []string
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				out = append(out, v1)
			}
		}
	}
	return out
}

func mapAllergenIngredient(foods []food) map[string][]string {
	allergens := make(map[string][]string)
	for _, v := range foods {
		for _, v1 := range v.allergen {
			a := allergens[v1]
			if a == nil {
				allergens[v1] = v.ingredient
			} else {
				allergens[v1] = intersection(a, v.ingredient)
			}
		}
	}
	return allergens
}

func allergenIngredients(a map[string][]string) map[string]bool {
	ingredients := make(map[string]bool)
	for _, v := range a {
		for _, v1 := range v {
			ingredients[v1] = true
		}
	}
	return ingredients
}

func part1(filename string) int {
	input := parseInput(filename)
	allergenIngredient := mapAllergenIngredient(input)
	ingredients := allergenIngredients(allergenIngredient)
	var answer int
	for _, v := range input {
		for _, v1 := range v.ingredient {
			if !ingredients[v1] {
				answer++
			}
		}
	}
	return answer
}

type allergenAndIngredient struct {
	allergen   string
	ingredient string
}

type allergenAndIngredientSlice []allergenAndIngredient

func (x allergenAndIngredientSlice) Len() int {
	return len(x)
}

func (x allergenAndIngredientSlice) Less(i, j int) bool {
	return x[i].allergen < x[j].allergen
}

func (x allergenAndIngredientSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func removeString(arr []string, s string) []string {
	for idx, v := range arr {
		if v == s {
			return append(arr[:idx], arr[idx+1:]...)
		}
	}
	return arr
}

func mapOneToOne(allergenIngredient map[string][]string) allergenAndIngredientSlice {
	var ai []allergenAndIngredient
	for len(ai) < len(allergenIngredient) {
		for k, v := range allergenIngredient {
			if len(v) == 1 {
				ai = append(ai, allergenAndIngredient{
					allergen:   k,
					ingredient: v[0],
				})
				for k1 := range allergenIngredient {
					allergenIngredient[k1] = removeString(allergenIngredient[k1], v[0])
				}
				break
			}
		}
	}
	return ai
}

func part2(filename string) string {
	input := parseInput(filename)
	allergenIngredient := mapAllergenIngredient(input)
	ai := mapOneToOne(allergenIngredient)
	sort.Sort(allergenAndIngredientSlice(ai))
	var s string
	for _, v := range ai {
		s += v.ingredient + ","
	}
	return s[:len(s)-1]
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 5 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != "mxmxvkd,sqjhc,fvjkl" {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
