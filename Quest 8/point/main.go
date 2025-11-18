package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func SetPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintInt(n int) { // recursive function
	if n < 0 {
		z01.PrintRune('-') // PrintInt(42)
		n = -n             // PrintInt(4) because 42/10=4
	} // Print 4
	if n >= 10 { // Print 2 42%10=2
		PrintInt(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	points := &point{} // points is a pointer pointing to a struct

	SetPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintInt(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintInt(points.y)
	z01.PrintRune('\n')
}
