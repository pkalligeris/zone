package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	first := true
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			// Print the first number (i) with two digits (leading zero if necessary)
			z01.PrintRune(rune(i/10 + '0')) // tens digit of i
			z01.PrintRune(rune(i%10 + '0')) // ones digit of i

			z01.PrintRune(' ') // space between the two numbers

			// Print the second number (j) with two digits (leading zero if necessary)
			z01.PrintRune(rune(j/10 + '0')) // tens digit of j
			z01.PrintRune(rune(j%10 + '0')) // ones digit of j

			first = false
		}
	}
	z01.PrintRune('\n')
}
