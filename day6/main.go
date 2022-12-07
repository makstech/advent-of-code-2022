package day6

import (
	"fmt"
	"os"
	"strings"
)

func Day6() {
	file, _ := os.ReadFile("day6/input.txt")
	input := string(file)
	list := strings.Split(input, "")
	var startOfPacket int
	var startOfPacket14 int

	for i := 4; i < len(list); i++ {
		if startOfPacket == 0 {
			chars := map[string]bool{}
			for _, c := range list[i-4 : i] {
				chars[c] = true
			}
			if len(chars) == 4 {
				startOfPacket = i
			}
		} else {
			chars14 := map[string]bool{}
			for _, c := range list[i-14 : i] {
				chars14[c] = true
			}
			if len(chars14) == 14 {
				startOfPacket14 = i
				break
			}
		}
	}

	fmt.Println("characters processed before the first start-of-packet marker is detected:", startOfPacket)
	fmt.Println("characters processed before the first start-of-message marker is detected:", startOfPacket14)
}
