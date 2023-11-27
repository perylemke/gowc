package cmd

import (
	"bufio"
	"os"
)

func CountLines(fileName string) (int, string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, "", err
	}

	return lines, file.Name(), nil
}
