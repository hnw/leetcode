package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	if l.Next == nil {
		return fmt.Sprintf("%d", l.Val)
	}
	return fmt.Sprintf("%d -> %s", l.Val, l.Next.String())
}

func addTwoNumbersAux(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if carry == 0 {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}
	}

	v := carry
	if l1 != nil {
		v += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		v += l2.Val
		l2 = l2.Next
	}
	return &ListNode{v % 10, addTwoNumbersAux(l1, l2, v/10)}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersAux(l1, l2, 0)
}

func main() {
	var (
		l1 *ListNode
		l2 *ListNode
	)
	l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	fmt.Println(addTwoNumbers(l1, l2)) // Output: 7 -> 0 -> 8
}
