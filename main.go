package main

import (
	"adventofcode/src/days"
	"adventofcode/src/utils"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	inputFilePath = "src/days/input_day%d.txt"
)

var (
	daysMap = map[int]utils.Day{
		1: &days.Day1{},
		2: &days.Day2{},
		3: &days.Day3{},
		4: &days.Day4{},
		5: &days.Day5{},
		6: &days.Day6{},
	}
	LineBreak = "\r\n"
)

func main() {
	initSolver()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("*---------------*")
	fmt.Println("AOC 2024 SOLVER")
	fmt.Println("*---------------*")
	fmt.Println("What day should we solve ?")
	dayInput, _ := reader.ReadString('\n')
	fmt.Println("Do you want to download the input file ? (y/n (or any other key to skip download))")
	dlInputFile, _ := reader.ReadString('\n')

	dayInput = strings.TrimSuffix(dayInput, LineBreak)
	dlInputFile = strings.TrimSuffix(dlInputFile, LineBreak)

	dayNumber, err := strconv.Atoi(dayInput)
	if err != nil {
		fmt.Printf("Invalid input, day should be a number between 1 and 24: %v", err)
		return
	}

	if dayNumber < 1 || dayNumber > 24 {
		fmt.Printf("Invalid input, day should be a number between 1 and 24: %v", err)
		return
	}

	fmt.Printf("Today we solve the day %d problem\n", dayNumber)
	day, ok := daysMap[dayNumber]
	if !ok {
		fmt.Printf("Day %d is not solved yet, maybe later...", dayNumber)
		return
	}

	filePath := fmt.Sprintf(inputFilePath, dayNumber)
	if dlInputFile == "y" {
		err = utils.DownloadInput(filePath, dayNumber)
		if err != nil {
			fmt.Printf("Cannot download input file for day %d: %v", dayNumber, err)
		}
	} else {
		fmt.Printf("Skip downloading input file for day %d\n", dayNumber)
	}

	solver := utils.DaySolver{}
	solver.Solve(day, filePath)
}

func initSolver() {
	if runtime.GOOS == "windows" {
		LineBreak = "\r\n"
	}
	if runtime.GOOS == "linux" {
		LineBreak = "\n"
	}
}
