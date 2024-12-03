package utils

import "fmt"

type Day interface {
	Preprocess() error
	Solve(path string)
	PartOneSolution() int
	PartTwoSolution() int
}

type DaySolver struct {
}

func (d *DaySolver) Solve(day Day, inputFilePath string) error {
	err := day.Preprocess()
	if err != nil {
		return err
	}

	day.Solve(inputFilePath)
	fmt.Printf("Problem has been solved: \nPt1 answer: %d\nPt2 answer: %d", day.PartOneSolution(), day.PartTwoSolution())
	return nil
}

type DayN struct {
	Pt1Sol int
	Pt2Sol int
}

func (d *DayN) Preprocess() error {
	return nil
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
