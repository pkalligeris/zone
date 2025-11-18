package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i, ch := range result {
		if ch >= 'A' && ch <= 'Z' {
			result[i] = ch - ('A' - 'a')
		}
	}
	return string(result)
}
