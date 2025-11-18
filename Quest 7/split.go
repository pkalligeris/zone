package piscine

// import "fmt"

func Split(s, sep string) []string {
	var result []string
	start := 0
	seplength := len(sep)

	if seplength == 0 { // split single character if separator is empty
		for _, char := range s {
			result = append(result, string(char))
		}
		return result
	}

	for i := 0; i <= len(s)-seplength; {
		if s[i:i+seplength] == sep {
			result = append(result, s[start:i])
			i += seplength
			start = i
		} else {
			i++
		}
	}

	// add remaining part after last separator
	result = append(result, s[start:])
	return result
}

/*func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
} */
