package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	file, _ := os.ReadFile("day10/input.txt")
	list := strings.Split(string(file), "\n")

	cycle := 0
	x := 1
	sumStrength := 0

	for _, s := range list {
		if s == "" {
			continue
		}
		prog := strings.Split(s, " ")
		cycles := 1
		if len(prog) == 2 {
			cycles = 2
		}
		for i := 0; i < cycles; i++ {
			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				sumStrength += cycle * x
			}
		}
		if cycles == 2 {
			addX, _ := strconv.Atoi(prog[1])
			x += addX
		}
	}

	fmt.Println("sum of the signal strengths:", sumStrength)
}
