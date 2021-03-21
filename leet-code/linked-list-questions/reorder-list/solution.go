/**
 * Definition for singly-linked list.
 */
package Solution

import (
	"container/list"
	"list_node"
)

func ReorderListWithStack(head *list_node.ListNode) {
	if head == nil {
		return
	}

	stack := list.New()
	fast := head.Next
	slow := head
    size := 1
	for fast != nil {
        size += 1
		stack.PushBack(slow)
		fast = fast.Next
		if fast != nil {
            size += 1
			fast = fast.Next
		}
		slow = slow.Next
	}

	var cur *list_node.ListNode
    stackTop := stack.Back()
	if size %2 != 0 {
		cur = slow
		slow = slow.Next
        cur.Next = nil
	}
	for stackTop != nil {
		next := slow.Next
		slow.Next = cur
		cur = stackTop.Value.(*list_node.ListNode)
        stackTop = stackTop.Prev()
		cur.Next = slow
		slow = next
	}
}

func reverseList(head *list_node.ListNode) *list_node.ListNode {
	cur := head
	var prev *list_node.ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}

func ReorderListWithoutSpace(head *list_node.ListNode) {
	if head == nil {
		return
	}

	fast := head
	slow := head
    size := 0
    var lastElementOfFirstHalf *list_node.ListNode
	for fast != nil {
        size += 1
		fast = fast.Next
		if fast != nil {
            size += 1
			fast = fast.Next
		}
        lastElementOfFirstHalf = slow
		slow = slow.Next
	}
    lastElementOfFirstHalf.Next = nil
	secondHalf := reverseList(slow)

	cur := head
	for secondHalf != nil {
		nextCur := cur.Next
		cur.Next = secondHalf
		nextInsecondHalf := secondHalf.Next
		cur = nextCur
		secondHalf.Next = cur
		secondHalf = nextInsecondHalf
	}
}
