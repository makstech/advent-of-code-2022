package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	file, _ := os.ReadFile("day1/input.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	maxCalories := 0
	currentCalories := 0
	var calories []int

	for _, s := range list {
		if s == "" {
			if currentCalories > 0 {
				calories = append(calories, currentCalories)
			}
			currentCalories = 0
		} else {
			n, _ := strconv.Atoi(s)
			currentCalories += n
		}

		if currentCalories > maxCalories {
			maxCalories = currentCalories
		}
	}

	fmt.Println("max calories: ", maxCalories)

	sort.Ints(calories)
	var totalCalories int
	for _, elfCalories := range calories[len(calories)-3:] {
		totalCalories += elfCalories
	}

	fmt.Println("total calories for top 3: ", totalCalories)
}
