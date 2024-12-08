package days

// pt2 8062916498078 too low

import (
	"adventofcode/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
	utils.DayN
}

var _ utils.Day = (*Day7)(nil)

func (d *Day7) Solve(path string) {
	lines := utils.ReadFile(path)
	total := 0
	for _, line := range lines {
		v, ok := solveLine(line)
		if ok {
			total += v
		}
	}
	d.Pt1Sol = total

	total = 0
	for _, line := range lines {
		v, ok := solveLine_pt2(line)
		if ok {
			total += v
		}
	}

	d.Pt2Sol = total
}

// Helper function to parse a line and solve it
func solveLine(line string) (int, bool) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return 0, false
	}

	// Parse target value
	target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, false
	}

	// Parse numbers
	numberStrings := strings.Fields(parts[1])
	nums := make([]int, len(numberStrings))
	for i, numStr := range numberStrings {
		nums[i], err = strconv.Atoi(numStr)
		if err != nil {
			return 0, false
		}
	}

	// Check if the equation is valid
	v, _ := evaluate(nums, target, 1, nums[0], fmt.Sprintf("%v", nums[0]))
	if v {
		// fmt.Printf("Solution: %v, %v\n", target, str)
		return target, true
	}
	return 0, false
}

// Helper function to parse a line and solve it
func solveLine_pt2(line string) (int, bool) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return 0, false
	}

	// Parse target value
	target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, false
	}

	// Parse numbers
	numberStrings := strings.Fields(parts[1])
	nums := make([]int, len(numberStrings))
	for i, numStr := range numberStrings {
		nums[i], err = strconv.Atoi(numStr)
		if err != nil {
			return 0, false
		}
	}

	// Check if the equation is valid
	v, _ := evaluate_pt2(nums, target, 1, nums[0], fmt.Sprintf("%v", nums[0]))
	if v {
		return target, true
	} else {
	}
	return 0, false
}

// Helper function to recursively evaluate the equation
func evaluate(nums []int, target int, index int, current int, expr string) (bool, string) {
	// Base case: If we've processed all numbers
	if index == len(nums) {
		return current == target, expr
	}

	// Recursively try both operators
	// Try addition
	newExpr := fmt.Sprintf("%s %s %d", expr, "+", nums[index])
	ok, v := evaluate(nums, target, index+1, current+nums[index], newExpr)
	if ok {
		return true, v
	}
	// Try multiplication
	newExpr = fmt.Sprintf("%s %s %d", expr, "*", nums[index])
	ok, v = evaluate(nums, target, index+1, current*nums[index], newExpr)
	if ok {
		return true, v
	}

	return false, ""
}

// Helper function to recursively evaluate the equation
func evaluate_pt2(nums []int, target int, index int, current int, expr string) (bool, string) {
	if index == len(nums) {
		return current == target, expr
	}

	// Try addition
	newExpr := fmt.Sprintf("%s %s %d", expr, "+", nums[index])
	ok, v := evaluate_pt2(nums, target, index+1, current+nums[index], newExpr)
	if ok {
		return true, v
	}

	// Try concatenation
	newExpr = fmt.Sprintf("%s %s %d", expr, "||", nums[index])
	val, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, nums[index]))
	ok, v = evaluate_pt2(nums, target, index+1, val, newExpr)
	if ok {
		return true, v
	}

	// Try multiplication
	newExpr = fmt.Sprintf("%s %s %d", expr, "*", nums[index])
	ok, v = evaluate_pt2(nums, target, index+1, current*nums[index], newExpr)
	if ok {
		return true, v
	}

	return false, ""
}
