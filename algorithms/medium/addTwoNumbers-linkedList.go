package main

import (
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// create and add a node to the list
func append(l *ListNode, val int) {
	l3 := &ListNode{Val: val}
	node := l
	for node.Next != nil {
		node = node.Next
	}
	node.Next = l3
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sl1, sl2 string

	// iterate over linked-lists and store values in reverse in a string
	for {
		sl1 = strconv.Itoa(l1.Val) + sl1

		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	for {
		sl2 = strconv.Itoa(l2.Val) + sl2

		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}

	//convert string to bigInt type for large numbers
	a, _ := new(big.Int).SetString(sl1, 0)
	b, _ := new(big.Int).SetString(sl2, 0)
	// sum values, convert to string
	a.Add(a, b)
	v := a.String()

	l3 := &ListNode{}

	// iterate over string in reverse and append each value to new list
	for i := len(v) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(v[i]))
		if i == len(v)-1 {
			l3 = &ListNode{Val: val}
		} else {
			append(l3, val)
		}

	}
	return l3
}
