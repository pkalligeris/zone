package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	length := max - min
	s := make([]int, length)

	for i := 0; i < length; i++ {
		s[i] = min + i
	}

	return s
}
