package day2

import (
	"fmt"
	"os"
	"strings"
)

var gMap = map[string]map[string]int{
	"A": {
		"X": 1 + 3,
		"Y": 2 + 6,
		"Z": 3 + 0,
	},
	"B": {
		"X": 1 + 0,
		"Y": 2 + 3,
		"Z": 3 + 6,
	},
	"C": {
		"X": 1 + 6,
		"Y": 2 + 0,
		"Z": 3 + 3,
	},
}

var actionMap = map[string]map[string]int{
	"A": {
		"X": 0 + 3,
		"Y": 3 + 1,
		"Z": 6 + 2,
	},
	"B": {
		"X": 0 + 1,
		"Y": 3 + 2,
		"Z": 6 + 3,
	},
	"C": {
		"X": 0 + 2,
		"Y": 3 + 3,
		"Z": 6 + 1,
	},
}

func Day2() {
	file, _ := os.ReadFile("day2/input.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	score := 0
	secondScore := 0

	for _, s := range list {
		if s == "" {
			continue
		}

		picks := strings.Split(s, " ")
		score += gMap[picks[0]][picks[1]]
		secondScore += actionMap[picks[0]][picks[1]]
	}

	fmt.Println("total score: ", score)
	fmt.Println("total second score: ", secondScore)
}
