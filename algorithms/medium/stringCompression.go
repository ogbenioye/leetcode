package main

import (
	"strconv"
)

func compress(chars []byte) int {
	var stChars = string(chars)
	var slByte []byte

	var sum = 1

	for i := 0; i < len(stChars); i++ {

		if i < len(stChars)-1 {

			if stChars[i] == stChars[i+1] {
				sum = sum + 1
			}
			if stChars[i] != stChars[i+1] {
				slByte = append(slByte, stChars[i])
				if sum > 1 {
					slByte = append(slByte, []byte(strconv.Itoa(sum))...)
				}
				sum = 1
			}
			if stChars[i] == stChars[i+1] && len(stChars)-1 == i+1 {
				slByte = append(slByte, stChars[i])
				slByte = append(slByte, []byte(strconv.Itoa(sum))...)
			}

			if stChars[i] != stChars[i+1] && len(stChars)-1 == i+1 {
				slByte = append(slByte, stChars[i+1])
			}

		} else if len(stChars) == 1 {
			slByte = append(slByte, stChars[i])
		}
	}
	return len(stChars)
}

func main() {
	data := []byte{'a', 'b'}
	compress(data)

}
