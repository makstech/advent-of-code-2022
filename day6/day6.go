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

	for i := 4; i < len(list); i++ {
		chars := map[string]bool{}
		for _, c := range list[i-4 : i] {
			chars[c] = true
		}
		if len(chars) == 4 {
			startOfPacket = i
			break
		}
	}

	fmt.Println("characters processed before the first start-of-packet marker is detected:", startOfPacket)
}
