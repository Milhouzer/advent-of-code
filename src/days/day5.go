package days

import (
	"adventofcode/src/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	utils.DayN
	before     map[int][]int
	startIndex int
}

var _ utils.Day = (*Day5)(nil)

var (
	orderingPattern = "%d|%d"
)

func (d *Day5) Preprocess(path string) error {
	d.before = make(map[int][]int)
	lines := utils.ReadFile(path)
	for i, line := range lines {
		if line == "" {
			d.startIndex = i
			break
		}

		var a int
		var b int
		fmt.Sscanf(line, orderingPattern, &a, &b)
		d.before[a] = append(d.before[a], b)
	}
	return nil
}

func (d *Day5) Solve(path string) {
	lines := utils.ReadFile(path)

	// Pt1
	total := 0
	for i := d.startIndex; i < len(lines); i++ {
		total += d.validLineMiddle(lines[i])
	}

	d.Pt1Sol = total

	// Pt2
	total = 0
	for j := d.startIndex; j < len(lines); j++ {
		total += d.invalidLineMiddle(lines[j])
	}

	d.Pt2Sol = total
}

func (d *Day5) invalidLineMiddle(line string) int {
	strVals := strings.Split(line, ",")

	var values []int
	for _, strVal := range strVals {
		v, _ := strconv.Atoi(strVal)
		values = append(values, v)
	}

	ordered, reordered := d.sort(values, true)

	if !reordered {
		return ordered[len(ordered)/2]
	}

	return 0
}

func (d *Day5) sort(values []int, firstTime bool) ([]int, bool) {
	reordered := false
	var ordered []int
	ordered = append(ordered, values...)

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if slices.Contains(d.before[values[j]], values[i]) {
				ordered[i], ordered[j] = ordered[j], ordered[i]
				reordered = true
			}
		}
	}

	if reordered {
		return d.sort(ordered, false)
	}

	return ordered, firstTime
}

func (d *Day5) validLineMiddle(line string) int {
	strVals := strings.Split(line, ",")

	var values []int
	for _, strVal := range strVals {
		v, _ := strconv.Atoi(strVal)
		values = append(values, v)
	}

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if slices.Contains(d.before[values[j]], values[i]) {
				return 0
			}
		}
	}

	return values[len(values)/2]
}
