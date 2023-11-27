package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

func main() {
	bytesFlag := flag.Bool("c", false, "print byte count")
	linesFlag := flag.Bool("l", false, "print line count")
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
}
