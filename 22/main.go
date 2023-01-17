package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open %q: %q\n", filename, err)
	}
	defer file.Close()

	var player1, player2 []int
	var second bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			second = true
			continue
		}
		if strings.HasPrefix(t, "Player") {
			continue
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalln(err)
		}
		if !second {
			player1 = append(player1, i)
		} else {
			player2 = append(player2, i)
		}
	}
	return player1, player2
}

func playToEnd(player1, player2 []int) []int {
	for len(player1) > 0 && len(player2) > 0 {
		if player1[0] > player2[0] {
			player1 = append(player1, player1[0], player2[0])
		} else {
			player2 = append(player2, player2[0], player1[0])
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}
	if len(player1) > 0 {
		return player1
	} else {
		return player2
	}
}

func calculateAnswer(deck []int) int {
	var answer int
	for i := range deck {
		answer += deck[i] * (len(deck) - i)
	}
	return answer
}

func part1(filename string) int {
	player1, player2 := parseInput(filename)
	deck := playToEnd(player1, player2)
	return calculateAnswer(deck)
}

func playRecursiveToEnd(player1, player2 []int) (bool, []int) {
	playHistory := make(map[string]bool)
	for len(player1) > 0 && len(player2) > 0 {
		s1 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(player1)), ","), "[]")
		s2 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(player2)), ","), "[]")
		if playHistory[s1+"|"+s2] {
			return true, nil
		}
		playHistory[s1+"|"+s2] = true
		var player1Won bool
		if player1[0] <= len(player1)-1 && player2[0] <= len(player2)-1 {
			newP1 := append([]int{}, player1[1:player1[0]+1]...)
			newP2 := append([]int{}, player2[1:player2[0]+1]...)
			player1Won, _ = playRecursiveToEnd(newP1, newP2)
		} else if player1[0] > player2[0] {
			player1Won = true
		}
		if player1Won {
			player1 = append(player1, player1[0], player2[0])
		} else {
			player2 = append(player2, player2[0], player1[0])
		}
		player1 = player1[1:]
		player2 = player2[1:]
	}
	if len(player1) > 0 {
		return true, player1
	} else {
		return false, player2
	}
}

func part2(filename string) int {
	player1, player2 := parseInput(filename)
	_, deck := playRecursiveToEnd(player1, player2)
	return calculateAnswer(deck)
}

func main() {
	testFile := "input_test.txt"
	if answer := part1(testFile); answer != 306 {
		log.Fatalln("Wrong answer, got:", answer)
	}
	if answer := part2(testFile); answer != 291 {
		log.Fatalln("Wrong answer, got", answer)
	}

	actualFile := "input.txt"
	fmt.Println(part1(actualFile))
	fmt.Println(part2(actualFile))
}
