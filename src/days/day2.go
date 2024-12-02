package days

import (
	"adventofcode/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day2 struct {
	utils.DayN
}

var _ utils.Day = (*Day2)(nil)

func (d *Day2) Solve(path string) {
	lines := utils.ReadFile(path)

	// Part 1
	safe := 0
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")

		// Report contains 0 or 1 level, it is safe
		if len(line) < 2 {
			safe++
			continue
		}

		// We could optimize this part and convert during calculations but this is really convenient
		levels := make([]int, len(line))
		var err error
		for j := 0; j < len(line); j++ {
			levels[j], err = strconv.Atoi(line[j])
			if err != nil {
				panic(fmt.Sprintf("invalid input %d: %v", i, line))
			}
		}

		// Get levels variation direction
		var sign float64 = 1
		v := levels[1] - levels[0]
		if v < 0 {
			sign = -1
		}

		// Check if level is actually safe
		isSafe := true
		for j := 0; j < len(levels)-1; j++ {
			variation := float64(levels[j+1] - levels[j])
			if variation*sign < 0 || math.Abs(variation) < 1 || math.Abs(variation) > 3 {
				isSafe = false
				break
			}
		}

		// Increment if safe
		if isSafe {
			safe++
		}
	}

	d.Pt1Sol = safe

	// Part 2 is the exact same code as the part 1 except instead of a boolean value [isSafe], we use an int value initialized at 1 and we remove 1
	// each time the report has an invalid property. If this value is negative at the end of the nested loop, it means the dampener can't make
	// the level safe
	safe = 0
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")

		// Report contains 0 or 1 level, it is safe
		if len(line) < 2 {
			safe++
			continue
		}

		// We could optimize this part and convert during calculations but this is really convenient
		levels := make([]int, len(line))
		var err error
		for j := 0; j < len(line); j++ {
			levels[j], err = strconv.Atoi(line[j])
			if err != nil {
				panic(fmt.Sprintf("invalid input %d: %v", i, line))
			}
		}

		// Get levels variation direction
		var sign float64 = 1
		v := levels[1] - levels[0]
		if v < 0 {
			sign = -1
		}

		// Check if level is actually safe
		isSafe := 1
		for j := 0; j < len(levels)-1; j++ {
			variation := float64(levels[j+1] - levels[j])
			if variation*sign < 0 || math.Abs(variation) < 1 || math.Abs(variation) > 3 {
				isSafe--
			}
		}

		// Increment if safe
		if isSafe >= 0 {
			safe++
		}
	}

	d.Pt2Sol = safe
}
