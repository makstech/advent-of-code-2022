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

func Day2() {
	file, _ := os.ReadFile("day2/input.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	score := 0

	for _, s := range list {
		if s == "" {
			continue
		}

		picks := strings.Split(s, " ")
		score += gMap[picks[0]][picks[1]]
	}

	fmt.Println("total score: ", score)
}
