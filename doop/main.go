package main

import "os"

func writeInt(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte{'0', '\n'})
		return
	}
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	os.Stdout.Write(buf[i:])
	os.Stdout.Write([]byte{'\n'})
}

func parseInt(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	i := 0
	neg := false
	if s[0] == '-' {
		neg = true
		i++
	}
	var n int64
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := int64(c - '0')
		if neg {
			if n < (-(1<<63)+d)/10 {
				return 0, false
			}
			n = n*10 - d
		} else {
			if n > ((1<<63-1)-d)/10 {
				return 0, false
			}
			n = n*10 + d
		}
	}
	return n, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := parseInt(os.Args[1])
	b, ok2 := parseInt(os.Args[3])
	if !ok1 || !ok2 {
		return
	}

	switch os.Args[2] {
	case "+":
		if (b > 0 && a > (1<<63-1)-b) || (b < 0 && a < -(1<<63)-b) {
			return
		}
		writeInt(a + b)
	case "-":
		if (b < 0 && a > (1<<63-1)+b) || (b > 0 && a < -(1<<63)+b) {
			return
		}
		writeInt(a - b)
	case "*":
		// Correct overflow check for multiplication
		if a > 0 {
			if b > 0 && a > (1<<63-1)/b {
				return
			}
			if b < 0 && b < -(1<<63)/a {
				return
			}
		} else if a < 0 {
			if b > 0 && a < -(1<<63)/b {
				return
			}
			if b < 0 && a != 0 && b < (1<<63-1)/a {
				return
			}
		}
		writeInt(a * b)
	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		writeInt(a / b)
	case "%":
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		writeInt(a % b)
	default:
		return
	}
}
