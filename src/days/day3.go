package days

import (
	"adventofcode/src/utils"
	"fmt"
	"regexp"
)

type Day3 struct {
	DayN
}

var _ Day = (*Day3)(nil)

var (
	pattern      = `mul\((\d{1,3}),(\d{1,3})\)|(don\'t\(\))|(do\(\))`
	doCmd        = "do()"
	dontCmd      = "don't()"
	mulPattern   = "mul(%d,%d)"
	solverRegexp *regexp.Regexp
)

// Preprocess regexp to avoid O(2^m) complexity addition to solve function. (m is the regexp length)
func (d *Day3) Preprocess(path string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	solverRegexp = re
	return nil
}

func (d *Day3) Solve(path string) {
	content := utils.ReadContent(path)

	sum := 0
	pt2sum := 0
	matches := solverRegexp.FindAllString(content, -1)
	factor := 1
	for _, match := range matches {
		if match == doCmd {
			factor = 1
		}

		if match == dontCmd {
			factor = 0
		}

		var a int
		var b int
		fmt.Sscanf(match, mulPattern, &a, &b)
		mul := a * b
		sum += mul
		pt2sum += mul * factor
	}

	d.Pt1Sol = sum
	d.Pt2Sol = pt2sum
}
