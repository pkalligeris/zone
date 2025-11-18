package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	if n < 1 || n > len(word) {
		return 0
	}
	return word[n-1]
}
