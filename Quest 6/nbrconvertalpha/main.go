package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	// No arguments → exit without printing anything
	if len(args) == 0 {
		return
	}

	// Check for uppercase flag
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
		// If only "--upper" was provided, exit without printing
		if len(args) == 0 {
			return
		}
	}

	for _, arg := range args {
		num := 0
		valid := true

		// Convert string to int manually
		for i := 0; i < len(arg); i++ {
			c := arg[i]
			if c < '0' || c > '9' {
				valid = false
				break
			}
			num = num*10 + int(c-'0')
		}

		// Print space for invalid or out-of-range numbers
		if !valid || num < 1 || num > 26 {
			z01.PrintRune(' ')
		} else {
			letter := rune('a' + num - 1)
			if upper {
				letter = rune('A' + num - 1)
			}
			z01.PrintRune(letter)
		}
	}

	z01.PrintRune('\n')
}
