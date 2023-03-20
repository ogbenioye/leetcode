package main

import (
	"fmt"
)

var digit []string
var palindrome []string

func isPalindrome(x int) bool {
	fmt.Println(x)
	num := fmt.Sprint(x)

	for i := 0; i < len(num); i++ {
		digit = append(digit, string(num[i]))
	}

	for j := len(num) - 1; j >= 0; j-- {
		palindrome = append(palindrome, string(num[j]))
	}

	for i, v := range digit {
		if v != palindrome[i] {
			return false
		}
	}
	fmt.Println(palindrome)
	fmt.Println(digit)
	return true
}

func main() {
	isPalindrome(123)
}
