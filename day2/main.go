package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type roundPoints struct {
	matchupPoint int
	choicePoint  int
}

const (
	rock     = iota + 1 // 1
	paper               // 2
	scissors            // 3
)

const (
	lose = iota     // 0
	tie  = iota + 2 // 2
	win  = iota + 4 // 6
)

func part1(opponent, yours string) int {

	roundOutcome := map[string]map[string]roundPoints{
		// rock
		"A": {
			"X": {tie, rock},      // tie - rock - 1
			"Y": {win, paper},     // win - paper - 2
			"Z": {lose, scissors}, // lose - scissors - 3
		},
		// paper
		"B": {
			"X": {lose, rock},    // lose - rock - 1
			"Y": {tie, paper},    // tie - paper - 2
			"Z": {win, scissors}, // win - scissors - 3
		},
		// scissors
		"C": {
			"X": {win, rock},     // win - rock - 1
			"Y": {lose, paper},   // lose - paper - 2
			"Z": {tie, scissors}, // tie - scissors - 3
		},
	}

	match := roundOutcome[opponent][yours]
	return calcPoints(match)

}

func part2(opponent, yours string) int {

	roundOutcome := map[string]map[string]roundPoints{
		// rock
		"A": {
			"X": {lose, scissors}, // lose - scissors - 3
			"Y": {tie, rock},      // tie - rock - 1
			"Z": {win, paper},     // win - paper - 2
		},
		// paper
		"B": {
			"X": {lose, rock},    // lose - rock - 1
			"Y": {tie, paper},    // tie - paper - 2
			"Z": {win, scissors}, // win - scissors - 3
		},
		// scissors
		"C": {
			"X": {lose, paper},   // lose - paper - 2
			"Y": {tie, scissors}, // tie - scissors - 3
			"Z": {win, rock},     // win - rock - 1
		},
	}
	match := roundOutcome[opponent][yours]
	return calcPoints(match)

}

func calcPoints(match roundPoints) int {
	return match.choicePoint + match.matchupPoint
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var counter_p1 int
	var counter_p2 int

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		points_p1 := part1(s[0], s[1])
		counter_p1 += points_p1

		points_p2 := part2(s[0], s[1])
		counter_p2 += points_p2

	}

	fmt.Printf("part 1: %v\n", counter_p1)
	fmt.Printf("part 2: %v\n", counter_p2)

}
