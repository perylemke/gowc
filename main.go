package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func countBytes(fileName string) (int, string, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return 0, "", err
	}

	return int(fileInfo.Size()), fileInfo.Name(), nil
}

func countLines(fileName string) (int, string, error) {
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

func countWords(fileName string) (int, string, error) {
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

func countRunes(fileName string) (int, string, error) {
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

func main() {
	bytesFlag := flag.Bool("c", false, "print byte count")
	linesFlag := flag.Bool("l", false, "print line count")
	wordsFlag := flag.Bool("w", false, "print word count")
	runeFlag := flag.Bool("m", false, "print rune count")
	flag.Parse()

	if *bytesFlag {
		size, name, err := countBytes(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", size, name)
	}

	if *linesFlag {
		lines, name, err := countLines(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", lines, name)
	}

	if *wordsFlag {
		words, name, err := countWords(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", words, name)
	}

	if *runeFlag {
		runes, name, err := countRunes(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", runes, name)

	}
}
