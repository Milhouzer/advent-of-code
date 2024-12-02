package utils

import "fmt"

type Day interface {
	Solve(path string)
	PartOneSolution() int
	PartTwoSolution() int
}

type DayN struct {
	Pt1Sol int
	Pt2Sol int
}

func (d *DayN) Solve(inputFile string) {
	fmt.Printf("not implemented")
}

func (d *DayN) PartOneSolution() int {
	return d.Pt1Sol
}

func (d *DayN) PartTwoSolution() int {
	return d.Pt2Sol
}

var _ Day = (*DayN)(nil)
