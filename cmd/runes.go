package cmd

import (
	"bufio"
	"os"
	"unicode/utf8"
)

func CountRunes(fileName string) (int, string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	runeCount := 0
	for scanner.Scan() {
		runeCount += utf8.RuneCountInString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, "", err
	}

	return runeCount, file.Name(), nil
}
