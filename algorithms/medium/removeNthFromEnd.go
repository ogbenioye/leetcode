package main

/*
	- Runtime: 2ms; Memory usage: 2.4MB (less than 15.06% of Go online submissions)
	- Faster than 60.15% of Go Online subbmissions as at 01-09-2023
	- Steps:
		Iterate and store values in each node in an array.
		Iterate over new array, remove the nth node and store the rest of the vales in a new array.
		Finally, add values from the last array into a new linked list.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := make([]int, 0)

	//store values from the list into an array
	for {
		arr = append(arr, head.Val)

		if head.Next == nil {
			break
		}

		head = head.Next
	}

	ar := make([]int, 0)

	//remove nth node
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-n {
			continue
		}
		ar = append(ar, arr[i])
	}

	if len(ar) == 0 {
		return nil
	}

	h := &ListNode{}

	//add values to new list
	for i := 0; i < len(ar); i++ {
		if i == 0 {
			h.Val = ar[i]
		} else {
			add(h, ar[i])
		}
	}

	return h
}

func add(head *ListNode, val int) {
	list := &ListNode{Val: val}

	for head.Next != nil {
		head = head.Next
	}

	head.Next = list
}
