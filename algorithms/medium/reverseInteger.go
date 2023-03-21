package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

func reverse(x int) int {
	var stringX []string
	sx := fmt.Sprint(x)

	for i := len(sx) - 1; i >= 0; i-- {
		stringX = append(stringX, string(sx[i]))
	}

	if math.Signbit(float64(x)) == true {
		stringX = append(stringX[:len(stringX)-1], stringX[len(stringX):]...)
		ns := strings.Join(stringX, "")
		na, _ := strconv.Atoi(ns)

		if bits.Len(uint(na)) > 32 || na > 2147483648 {
			return 0
		}

		ns1 := "-" + ns
		nx, _ := strconv.Atoi(ns1)
		return nx
	} else {
		ns := strings.Join(stringX, "")
		nx, _ := strconv.Atoi(ns)

		if bits.Len(uint(nx)) > 32 || nx > 2147483647 {
			return 0
		}
		return nx
	}
}

func main() {
	sm := 2147483651
	reverse(sm)
}
