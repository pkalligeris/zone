package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(filename string) {
	msg := "ERROR: open " + filename + ": no such file or directory\n"
	for _, r := range msg {
		z01.PrintRune(r)
	}
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		printError(filename)
		os.Exit(1)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			for _, b := range buf[:n] {
				z01.PrintRune(rune(b))
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			printError(filename)
			os.Exit(1)
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Read from stdin
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				for _, b := range buf[:n] {
					z01.PrintRune(rune(b))
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				os.Exit(1)
			}
		}
	} else {
		for _, filename := range args {
			printFile(filename)
		}
	}
}
