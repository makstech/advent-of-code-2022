package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() {
	file, _ := os.ReadFile("day4/input.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	pairsContain := 0
	pairsOverlap := 0

	for _, s := range list {
		if s == "" {
			continue
		}

		pair := strings.Split(s, ",")
		elf1 := strings.Split(pair[0], "-")
		elf1From, _ := strconv.Atoi(elf1[0])
		elf1To, _ := strconv.Atoi(elf1[1])
		elf2 := strings.Split(pair[1], "-")
		elf2From, _ := strconv.Atoi(elf2[0])
		elf2To, _ := strconv.Atoi(elf2[1])

		if elf2From >= elf1From && elf2To <= elf1To || elf1From >= elf2From && elf1To <= elf2To {
			pairsContain++
		}

		if elf2From >= elf1From && elf2From <= elf1To || elf2To <= elf1To && elf2To >= elf1From ||
			elf1From >= elf2From && elf1From <= elf2To || elf1To <= elf2To && elf1To >= elf2From {
			pairsOverlap++
		}
	}

	fmt.Println("pairs fully contain the other:", pairsContain)
	fmt.Println("pairs overlap:", pairsOverlap)
}
