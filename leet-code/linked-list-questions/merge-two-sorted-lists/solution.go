package Solution

import (
	"list_node"
)

func MergeTwoSortedListsWithExtraSpace(l1 *list_node.ListNode, l2 *list_node.ListNode) *list_node.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newListHead := &list_node.ListNode{
		Val: 0,
	}
	if l1.Val < l2.Val {
		newListHead.Val = l1.Val
		l1 = l1.Next
	} else {
		newListHead.Val = l2.Val
		l2 = l2.Next
	}
	cur := newListHead

	for l1 != nil && l2 != nil {
		cur.Next = &list_node.ListNode{
			Val: 0,
		}
		if l1.Val < l2.Val {
			cur.Next.Val = l1.Val
			l1 = l1.Next
		} else {
			cur.Next.Val = l2.Val
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return newListHead
}

func MergeTwoSortedListsInPlace(l1 *list_node.ListNode, l2 *list_node.ListNode) *list_node.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newListHead *list_node.ListNode
	if l1.Val < l2.Val {
		newListHead = l1
		l1 = l1.Next
	} else {
		newListHead = l2
		l2 = l2.Next
	}

	prev := newListHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			prev = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			prev = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		prev.Next = l1
	} else if l2 != nil {
		prev.Next = l2
	}

	return newListHead
}
