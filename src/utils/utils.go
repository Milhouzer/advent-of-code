package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func ReadFile(filename string) []string {
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

func ReadFileContent(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
