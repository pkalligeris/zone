package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	number := []int{}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		number = append(number, n%10)
		n /= 10
	}
	for i := 0; i < len(number); i++ {
		for j := i + 1; j < len(number); j++ {
			if number[j] < number[i] {
				number[i], number[j] = number[j], number[i]
			}
		}
	}
	for _, k := range number {
		z01.PrintRune(rune('0' + k))
	}
}
