package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	name := ""

	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' {
			name = programName[i+1:]
			break
		}
		if i == 0 {
			name = programName
		}
	}

	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
