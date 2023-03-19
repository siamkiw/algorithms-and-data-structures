package listleet

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1 := list1
	l2 := list2
	var current *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	for l1 != nil || l2 != nil {
		fmt.Println("current :", current)
		if l1 != nil && l2 == nil {
			current.Next = l1.Next
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			current.Next = l2.Next
			l2 = l2.Next
		} else if l1.Val > l2.Val {
			current.Next = l1
			current = l1
			current.Next = l2
			current = l2
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Val < l2.Val {
			current.Next = l2
			current = l2
			current.Next = l1
			current = l1
			l1 = l1.Next
			l2 = l2.Next
		}

	}
	return current
}
