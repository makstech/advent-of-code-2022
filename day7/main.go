package day7

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day7() {
	file, _ := os.ReadFile("day7/input.txt")
	input := string(file)
	list := strings.Split(input, "\n")

	totalSize100k := 0
	commandRegexp := regexp.MustCompile(`\$ (\w+) ?(.+)?`)
	lsRegexp := regexp.MustCompile(`(dir|\d+) (.+)`)
	curPos := ""
	sizeMap := map[string]int{}

	for _, s := range list {
		if s == "" {
			continue
		}

		if s[:1] == "$" {
			command := commandRegexp.FindStringSubmatch(s)
			dir := command[2]
			if command[1] == "cd" {
				if dir == "/" {
					continue
				} else if dir == ".." {
					curPos = curPos[:strings.LastIndex(curPos, "_")]
				} else {
					curPos += "_" + dir
				}
			}
		} else {
			ls := lsRegexp.FindStringSubmatch(s)
			if ls[1] != "dir" {
				size, _ := strconv.Atoi(ls[1])
				positions := strings.Split(curPos, "_")

				for i := 1; i < len(positions)+1; i++ {
					sizeMap[strings.Join(positions[0:i], "_")] += size
				}
			}
		}
	}

	needMoreFree := -70000000 + 30000000 + sizeMap[""]
	dirToDeleteSize := sizeMap[""]

	for _, s := range sizeMap {
		if s <= 100_000 {
			totalSize100k += s
		}
		if s >= needMoreFree && s < dirToDeleteSize {
			dirToDeleteSize = s
		}
	}

	fmt.Println("total size of at most 100000:", totalSize100k)
	fmt.Println("total size of directory to delete:", dirToDeleteSize)
}
