package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	digit := false
	next := 0

	for _, c := range s {
		if c == '-' && !digit {
			next = -1
		} else if c >= '0' && c <= '9' {
			if !digit && next != 0 {
				sign = next
			}
			digit = true
			num = num*10 + int(c-'0')
		}
	}

	if !digit {
		return 0
	}

	return num * sign
}
