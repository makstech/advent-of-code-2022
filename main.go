package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day11"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Bₒₒₜᵢₙg ₐₗₗ ₜₕₑ ᵤₙₙₑcₑₛₛₐᵣy cₒₛₘᵢc ᵣₐdᵢₐₜᵢₒₙ fₒᵣ ₜₕᵢₛ ₛₕᵢₜcₒdₑ ₜₒ wₒᵣₖ")
	fmt.Print("ₚᵢcₖ ₐ dₐy ... ")
	var day string
	fmt.Scanln(&day)

	funcs := map[string]func(){
		"1":  day1.Day1,
		"2":  day2.Day2,
		"3":  day3.Day3,
		"4":  day4.Day4,
		"5":  day5.Day5,
		"6":  day6.Day6,
		"7":  day7.Day7,
		"8":  day8.Day8,
		"9":  day9.Day9,
		"10": day10.Day10,
		"11": day11.Day11,
	}

	if fun, ok := funcs[day]; ok {
		fmt.Println("ₛₜₐᵣₜᵢₙg dₐy ...", day)
		fun()
	} else {
		log.Fatal("day doesn't exist")
	}
}
