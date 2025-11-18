package piscine

func ForEach(f func(int), a []int) {
	for _, numb := range a {
		f(numb)
	}
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	ForEach(PrintNbr, a)
// }
