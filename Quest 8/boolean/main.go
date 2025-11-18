package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	return nbr%2 == 0 // isEven returns true if number is even, else false
}

func main() {
	lengttofArg := len(os.Args[1:]) // len(os.Args) gives how many arguments there are.
	// Without the programName
	evenMsg := "I have an even number of arguments"
	oddMsg := "I have an odd number of arguments"

	if isEven(lengttofArg) {
		printStr(evenMsg)
	} else {
		printStr(oddMsg)
	}
}
