package cmd

import (
	"bufio"
	"os"
	"strings"
)

func CountWords(fileName string) (int, string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		return 0, "", err
	}

	return wordCount, file.Name(), nil
}
