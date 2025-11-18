package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if nb == 0 {
		return 0
	}

	result := 1
	for i := 0; i <= power-1; i++ {
		result = result * nb
	}
	return result
}
