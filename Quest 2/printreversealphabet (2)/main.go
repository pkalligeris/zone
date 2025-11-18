package main

import "github.com/01-edu/z01"

func main() {
	for n := 'z'; n >= 'a'; n-- {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
