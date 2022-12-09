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
	secondKnotVisited := map[string]bool{}
	rope := make([][2]int, 10)
	for i := 0; i < 10; i++ {
		rope[i] = [2]int{0, 0}
	}

	for _, s := range list {
		if s == "" {
			continue
		}
		moves := strings.Split(s, " ")
		steps, _ := strconv.Atoi(moves[1])
		for i := 0; i < steps; i++ {
			if moves[0] == "U" {
				rope[0][1]++
			} else if moves[0] == "D" {
				rope[0][1]--
			} else if moves[0] == "R" {
				rope[0][0]++
			} else if moves[0] == "L" {
				rope[0][0]--
			}
			for k := 1; k < 10; k++ {
				xDif := rope[k-1][0] - rope[k][0]
				yDif := rope[k-1][1] - rope[k][1]
				if Abs(xDif) < 2 && Abs(yDif) < 2 {
					break
				}
				if xDif == 0 {
					if yDif > 0 {
						rope[k][1]++
					} else {
						rope[k][1]--
					}
					continue
				}
				if yDif == 0 {
					if xDif > 0 {
						rope[k][0]++
					} else {
						rope[k][0]--
					}
					continue
				}
				if xDif > 0 {
					rope[k][0]++
				} else {
					rope[k][0]--
				}
				if yDif > 0 {
					rope[k][1]++
				} else {
					rope[k][1]--
				}
			}
			secondKnotVisited[strconv.Itoa(rope[1][0])+":"+strconv.Itoa(rope[1][1])] = true
			tailVisited[strconv.Itoa(rope[9][0])+":"+strconv.Itoa(rope[9][1])] = true
		}
	}

	fmt.Println("positions 2nd knot of the rope visited:", len(secondKnotVisited))
	fmt.Println("positions tail of the rope visited:", len(tailVisited))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
