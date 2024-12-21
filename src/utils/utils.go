package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

const (
	baseUrl = "https://adventofcode.com/2024/day/%d/input"
)

func DownloadInput(outputPath string, day int) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cookie := os.Getenv("AOC_SESSION")
	if cookie == "" {
		return fmt.Errorf("AOC_SESSION environment variable is not set")
	}

	url := fmt.Sprintf(baseUrl, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", cookie))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Input for day %d downloaded successfully to %s", day, outputPath)
	return nil
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadLinesInPacks(filename string, packSize int) [][]string {
	if packSize <= 0 {
		log.Fatal("Pack size must be greater than 0")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var packs [][]string
	var currentPack []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentPack = append(currentPack, scanner.Text())
		if len(currentPack) == packSize {
			packs = append(packs, currentPack)
			currentPack = nil // Start a new pack
		}
	}

	// Add any remaining lines that don't form a full pack
	if len(currentPack) > 0 {
		packs = append(packs, currentPack)
	}

	return packs
}

func ReadContent(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

// ClearScreen clears the console screen using an ANSI escape sequence
func ClearScreen() {
	cmd := exec.Command("clear")
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func IsFloatAnInt(f float64) bool {
	if f == float64(int64(f)) {
		return true
	}

	const epsilon = 1e-9 // Tolerance for floating-point comparisons
	return math.Abs(f-math.Round(f)) < epsilon
}
