package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	countneg := 0
	countpos := 0

	for i := 0; i < len(a)-1; i++ {
		val := f(a[i], a[i+1])

		if val <= 0 {
			countneg++
		}
		if val >= 0 {
			countpos++
		}
	}

	// If all pairs are <=0 → ascending
	if countneg == len(a)-1 {
		return true
	}
	// If all pairs are >=0 → descending
	if countpos == len(a)-1 {
		return true
	}

	// Otherwise, not sorted
	return false
}

// func f(a, b int) int {
// 	return a - b
// }

// func main() {
// 	a1 := []int{0, 1, 2, 3, 4, 5}
// 	a2 := []int{0, 2, 1, 3}

// 	result1 := IsSorted(f, a1)
// 	result2 := IsSorted(f, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
