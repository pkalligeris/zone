package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i, word := range a {
		for _, r := range word {
			z01.PrintRune(r)
		}
		if i != len(a)-1 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}

/*func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}*/
