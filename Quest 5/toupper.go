package piscine

func ToUpper(s string) string {
	result := []rune(s)
	for i, ch := range result {
		if ch >= 'a' && ch <= 'z' {
			result[i] = ch - ('a' - 'A')
		}
	}
	return string(result)
}
