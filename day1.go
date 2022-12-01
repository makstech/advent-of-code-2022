package main

import (
	"os"
	"strconv"
	"strings"
)

func Day1() {
	file, _ := os.ReadFile("input/day1.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	maxCalories := 0
	currentCalories := 0

	for _, s := range list {
		if s == "" {
			currentCalories = 0
		} else {
			n, _ := strconv.Atoi(s)
			currentCalories += n
		}

		if currentCalories > maxCalories {
			maxCalories = currentCalories
		}
	}

	println("max calories: ", maxCalories)
}
