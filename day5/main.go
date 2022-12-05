package day5

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day5() {
	file, _ := os.ReadFile("day5/input.txt")
	input := string(file)
	parts := strings.Split(input, "\n\n")

	textStacks := regexp.MustCompile("[1-9]").FindAllString(parts[0], -1)
	totalStacks := len(textStacks)
	textCrates := regexp.MustCompile("[ \\[]([A-Z ])[ \\]][ \\n]").FindAllStringSubmatch(parts[0], -1)
	stacks := map[string][]string{}

	for i, c := range textCrates {
		if c[1] == " " {
			continue
		}

		stacks[textStacks[i%totalStacks]] = append(stacks[textStacks[i%totalStacks]], c[1])
	}

	actions := regexp.MustCompile("move (\\d+) from (\\d) to (\\d)").FindAllStringSubmatch(parts[1], -1)
	for _, a := range actions {
		count, _ := strconv.Atoi(a[1])
		var crates []string
		for i := count - 1; i >= 0; i-- {
			crates = append(crates, stacks[a[2]][i])
		}
		stacks[a[3]] = append(crates, stacks[a[3]]...)
		stacks[a[2]] = stacks[a[2]][count:]
	}

	topCrates := ""
	for _, s := range textStacks {
		topCrates += stacks[s][0]
	}

	fmt.Println("top crates:", topCrates)
}
