package day11

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day11() {
	file, _ := os.ReadFile("day11/input.txt")
	list := strings.Split(string(file), "\n\n")
	type Monkey struct {
		items     []int
		operation struct {
			action string
			second int
			old    bool
		}
		division  int
		throw     map[bool]int
		inspected int
	}
	var monkeys []Monkey
	num := regexp.MustCompile("[0-9]+")
	for _, m := range list {
		monkey := Monkey{items: []int{}, throw: map[bool]int{}}
		rows := strings.Split(m, "\n")
		for _, itemS := range num.FindAllString(rows[1], -1) {
			item, _ := strconv.Atoi(itemS)
			monkey.items = append(monkey.items, item)
		}
		var second string
		fmt.Sscanf(strings.Split(rows[2], " = ")[1], "old %s %s", &monkey.operation.action, &second)
		if second == "old" {
			monkey.operation.old = true
		} else {
			secondI, _ := strconv.Atoi(second)
			monkey.operation.second = secondI
		}
		fmt.Sscanf(rows[3], "  Test: divisible by %d", &monkey.division)
		var to int
		fmt.Sscanf(rows[4], "    If true: throw to monkey %d", &to)
		monkey.throw[true] = to
		fmt.Sscanf(rows[5], "    If false: throw to monkey %d", &to)
		monkey.throw[false] = to
		monkeys = append(monkeys, monkey)
	}
	for _, monkey := range monkeys {
		fmt.Println(monkey)
	}
	for r := 0; r < 20; r++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for _, item := range monkey.items {
				second := monkey.operation.second
				if monkey.operation.old {
					second = item
				}
				var stress int
				switch monkey.operation.action {
				case "+":
					stress = item + second
				case "*":
					stress = item * second
				}
				stress = stress / 3
				throwTo := monkey.throw[stress%monkey.division == 0]
				monkeys[throwTo].items = append(monkeys[throwTo].items, stress)
			}
			monkey.inspected += len(monkey.items)
			monkey.items = []int{}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	fmt.Println("level of monkey business after 20 rounds of stuff-slinging simian shenanigans:", monkeys[0].inspected*monkeys[1].inspected)
}
