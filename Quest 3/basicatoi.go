package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, i := range s {
		result = result*10 + int(i-'0')
	}
	return result
}
