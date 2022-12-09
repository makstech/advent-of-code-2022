package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9() {
	file, _ := os.ReadFile("day9/input.txt")
	list := strings.Split(string(file), "\n")

	tailVisited := map[string]bool{}
	tailPos := []int{0, 0}
	headPos := []int{0, 0}

	for _, s := range list {
		if s == "" {
			continue
		}
		moves := strings.Split(s, " ")
		steps, _ := strconv.Atoi(moves[1])
		for i := 0; i < steps; i++ {
			xDir := 0
			yDir := 0
			if moves[0] == "U" {
				headPos[1]++
				yDir++
			} else if moves[0] == "D" {
				headPos[1]--
				yDir--
			} else if moves[0] == "R" {
				headPos[0]++
				xDir++
			} else if moves[0] == "L" {
				headPos[0]--
				xDir--
			}
			xDif := headPos[0] - tailPos[0]
			yDif := headPos[1] - tailPos[1]
			if strings.Contains("LR", moves[0]) {
				if Abs(xDif) > 1 {
					tailPos[0] = tailPos[0] + xDir
					if Abs(yDif) > 0 {
						tailPos[1] = tailPos[1] + yDif
					}
				}
			} else {
				if Abs(yDif) > 1 {
					tailPos[1] = tailPos[1] + yDir
					if Abs(xDif) > 0 {
						tailPos[0] = tailPos[0] + xDif
					}
				}
			}
			tailVisited[strconv.Itoa(tailPos[0])+":"+strconv.Itoa(tailPos[1])] = true
		}
	}

	fmt.Println("positions tail of the rope visit:", len(tailVisited))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
