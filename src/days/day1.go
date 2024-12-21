package days

import (
	"adventofcode/src/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct {
	DayN
}

var _ Day = (*Day1)(nil)

func (d *Day1) Solve(inputFile string) {
	lines := utils.ReadLines(inputFile)

	// Part 1
	sum := 0
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	rOccurences := make(map[int]int)
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], "   ")
		lv, err := strconv.Atoi(strings.TrimSpace(line[0]))
		if err != nil {
			panic(fmt.Sprintf("invalid input: %v", lines[i]))
		}
		rv, err := strconv.Atoi(strings.TrimSpace(line[1]))
		if err != nil {
			panic(fmt.Sprintf("invalid input: %v", lines[i]))
		}

		left[i] = lv
		right[i] = rv
		_, ok := rOccurences[rv]
		if ok {
			rOccurences[rv]++
		} else {
			rOccurences[rv] = 1
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(right[i] - left[i])))
	}

	d.Pt1Sol = sum

	// Part 2
	for i := 0; i < len(left); i++ {
		v, ok := rOccurences[int(left[i])]
		if ok {
			sum += v * int(left[i])
		}
	}

	d.Pt2Sol = sum
}
