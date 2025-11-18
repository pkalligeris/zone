package piscine

func Capitalize(s string) string {
	b := []byte(s)
	word := false

	for i := 0; i < len(b); i++ {
		c := b[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			if !word && c >= 'a' && c <= 'z' { // first letter of word
				b[i] = c - 'a' + 'A'
			} else if word && c >= 'A' && c <= 'Z' { // not first letter
				b[i] = c - 'A' + 'a'
			}
			word = true
		} else {
			word = false
		}
	}

	return string(b)
}
