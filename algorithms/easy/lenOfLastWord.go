package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	res := strings.Fields(s)
	return len(res[len(res)-1])
}

func main() {
	str := "   fly me   to   the moon  "
	lengthOfLastWord(str)
}
