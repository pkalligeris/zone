package piscine

func SplitWhiteSpaces(s string) []string {
	var word []string
	var currentword []rune

	for i := 0; i < len(s); i++ { // check for space and newline
		if s[i] == ' ' || s[i] == '\n' {
			if len(currentword) > 0 { // if we find word append and reset word
				word = append(word, string(currentword))
				currentword = nil

			}
		} else { // else add char to word
			currentword = append(currentword, rune(s[i]))
		}
	}

	if len(currentword) > 0 { // for the last word because it will never find '\n'
		word = append(word, string(currentword))
	}
	return word
}

/*func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
} */
