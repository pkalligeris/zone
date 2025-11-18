package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 0 || n >= 10 {
		return
	}
	combRec([]int{}, 0, n)
	z01.PrintRune('\n')
}

func combRec(current []int, start, n int) {
	if len(current) == n {
		for _, d := range current {
			z01.PrintRune(rune('0' + d))
		}

		if current[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 9; i++ {
		combRec(append(current, i), i+1, n)
	}
}
