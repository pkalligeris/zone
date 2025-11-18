package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == '+' {
		i = 1
	}

	result := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}

	return result * sign
}
