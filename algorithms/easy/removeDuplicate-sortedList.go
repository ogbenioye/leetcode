package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := head

	if h == nil {
		return h
	}

	for h.Next != nil {
		if h.Val == h.Next.Val {
			h = delete(h, h.Val)
		} else {
			h = h.Next
		}

		if h.Next == nil {
			break
		}
	}
	return head
}

func delete(l *ListNode, val int) *ListNode {
	toDelete := l
	for toDelete.Next != nil {
		if val == toDelete.Next.Val {
			toDelete.Next = toDelete.Next.Next
			return l
		}
		toDelete = toDelete.Next
	}
	return l
}
