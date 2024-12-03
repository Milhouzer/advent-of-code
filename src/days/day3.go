package days

import (
	"adventofcode/src/utils"
	"fmt"
	"regexp"
)

type Day3 struct {
	utils.DayN
}

var _ utils.Day = (*Day3)(nil)

var (
	pattern    = `mul\((\d{1,3}),(\d{1,3})\)|(don\'t\(\))|(do\(\))`
	doCmd      = "do()"
	dontCmd    = "don't()"
	mulPattern = "mul(%d,%d)"
)

func (d *Day3) Solve(path string) {
	content := utils.ReadFileContent(path)

	sum := 0
	pt2sum := 0
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(content, -1)
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