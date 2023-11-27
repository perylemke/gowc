package main

import (
	"flag"
	"fmt"

	"github.com/perylemke/gowc/cmd"
)

func main() {
	// Flags
	bytesFlag := flag.Bool("c", false, "print byte count")
	linesFlag := flag.Bool("l", false, "print line count")
	wordsFlag := flag.Bool("w", false, "print word count")
	runeFlag := flag.Bool("m", false, "print rune count")
	flag.Parse()

	if *bytesFlag {
		size, name, err := cmd.CountBytes(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", size, name)
	}

	if *linesFlag {
		lines, name, err := cmd.CountLines(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", lines, name)
	}

	if *wordsFlag {
		words, name, err := cmd.CountWords(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", words, name)
	}

	if *runeFlag {
		runes, name, err := cmd.CountRunes(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("  %d %s\n", runes, name)
	}

	if !*bytesFlag && !*linesFlag && !*wordsFlag && !*runeFlag {
		bytes, name, err := cmd.CountBytes(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}

		lines, _, err := cmd.CountLines(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}

		words, _, err := cmd.CountWords(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("  %d %d %d %s\n", lines, words, bytes, name)
	}
}
