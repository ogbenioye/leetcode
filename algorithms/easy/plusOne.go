package main

import (
	"math/big"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	in := sliceToInt(digits)

	one := big.NewInt(1)
	addOne := in.Add(in, one)

	newDigits := splitInt(addOne)

	return newDigits
}

func sliceToInt(s []int) *big.Int {
	var strs []string

	for _, v := range s {
		str := strconv.Itoa(v)
		strs = append(strs, str)
	}

	st := strings.Join(strs, "")
	num, _ := new(big.Int).SetString(st, 0)
	return num
}

func splitInt(n *big.Int) []int {
	newN := n.String()
	var sl []int
	for i := 0; i < len(newN); i++ {
		newV, _ := strconv.Atoi(string(newN[i]))
		sl = append(sl, newV)
	}
	return sl
}

func main() {
	arr := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	plusOne(arr)
}
