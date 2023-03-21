package main

import (
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	var ns string
	s = strings.TrimSpace(s)

	for i := 0; i < len(s); i++ {
		t, _ := strconv.Atoi(string(s[i]))
		if t == 0 && string(s[i]) != "0" && string(s[i]) != "-" && string(s[i]) != "+" {
			break
		}

		if string(s[i]) == "-" && i != 0 {
			break
		}

		if string(s[i]) == "+" && i != 0 {
			break
		}
		ns = ns + string(s[i])
	}
	d, _ := strconv.Atoi(ns)

	if d > 2147483647 {
		d = 2147483647
	} else if d < -2147483648 {
		d = -2147483648
	}
	return d
}

func main() {
	myAtoi("42")
}
