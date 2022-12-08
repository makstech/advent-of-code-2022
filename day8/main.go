package day8

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day8() {
	file, _ := os.ReadFile("day8/input.txt")
	list := strings.Split(string(file), "\n")

	rows := len(list)
	columns := len(list[0])
	grid := make([][]int, rows)
	gridColumns := make([][]int, columns)
	for i, r := range list {
		if r == "" {
			continue
		}

		grid[i] = make([]int, columns)
		for ii, t := range strings.Split(r, "") {
			if i == 0 {
				gridColumns[ii] = make([]int, rows)
			}

			tree, _ := strconv.Atoi(t)
			grid[i][ii] = tree
			gridColumns[ii][i] = tree
		}
	}

	visibleTrees := 0
	for i, r := range grid {
		rowHighest := 0
		for ii, t := range r {
			if t > rowHighest {
				rowHighest = t
				visibleTrees++
				continue
			}
			if i == 0 || ii == 0 || i == rows-1 || ii == columns-1 {
				visibleTrees++
				continue
			}
			tressFrontS := r[ii+1:]
			tressFront := make([]int, len(tressFrontS))
			copy(tressFront, tressFrontS)
			sort.Ints(tressFront)
			if t > tressFront[len(tressFront)-1] {
				visibleTrees++
				continue
			}
			treesAboveS := gridColumns[ii][:i]
			treesAbove := make([]int, len(treesAboveS))
			copy(treesAbove, treesAboveS)
			sort.Ints(treesAbove)
			if t > treesAbove[i-1] {
				visibleTrees++
				continue
			}
			treesBelowS := gridColumns[ii][i+1:]
			treesBelow := make([]int, len(treesBelowS))
			copy(treesBelow, treesBelowS)
			sort.Ints(treesBelow)
			if t > treesBelow[len(treesBelow)-1] {
				visibleTrees++
				continue
			}
		}
	}

	fmt.Println("visible trees", visibleTrees)
}
