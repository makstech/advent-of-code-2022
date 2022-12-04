package day3

import (
	"fmt"
	"os"
	"strings"
)

func Day3() {
	file, _ := os.ReadFile("day3/input.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	priorities := map[string]int{}
	prio := 1
	for i := 'a'; i <= 'z'; i++ {
		priorities[string(i)] = prio
		prio++
	}
	for i := 'A'; i <= 'Z'; i++ {
		priorities[string(i)] = prio
		prio++
	}

	totalPriority := 0

	for _, s := range list {
		if s == "" {
			continue
		}

		compartment := len(s) / 2
		totalPriority += priorities[FindSame(s[:compartment], s[compartment:])]
	}

	fmt.Println("total priority:", totalPriority)

	totalGroupPriority := 0
	for i := 0; i < len(list); i += 3 {
		for _, c := range list[i] {
			letter := string(c)
			if strings.Contains(list[i+1], letter) && strings.Contains(list[i+2], letter) {
				totalGroupPriority += priorities[letter]
				break
			}
		}
	}

	fmt.Println("total group priority:", totalGroupPriority)
}

func FindSame(a, b string) string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range strings.Split(b, "") {
		mb[x] = struct{}{}
	}

	for _, x := range strings.Split(a, "") {
		if _, found := mb[x]; found {
			return x
		}
	}

	return ""
}
