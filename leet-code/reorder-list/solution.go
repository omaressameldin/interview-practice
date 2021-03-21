/**
 * Definition for singly-linked list.
 */
package Solution

import (
	"container/list"
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateFromSlice(items []int)  *ListNode {
	if len(items) == 0 {
		return nil
	}

	head := &ListNode{
		Val: items[0],
		Next: nil,
	}
	cur := head
	for _, item := range items[1:] {
		cur.Next = &ListNode{
			Val: item,
			Next: nil,
		}
		cur = cur.Next
	}

	return head
}

func (list *ListNode) String() string {
	var b strings.Builder
	cur := list
	for cur != nil {
		b.WriteString(fmt.Sprintf("%d", cur.Val))
		cur = cur.Next
	}

	return b.String()
}

func ReorderListWithStack(head *ListNode) {
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

	var cur *ListNode
    stackTop := stack.Back()
	if size %2 != 0 {
		cur = slow
		slow = slow.Next
        cur.Next = nil
	}
	for stackTop != nil {
		next := slow.Next
		slow.Next = cur
		cur = stackTop.Value.(*ListNode)
        stackTop = stackTop.Prev()
		cur.Next = slow
		slow = next
	}
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}

func ReorderListWithoutSpace(head *ListNode) {
	if head == nil {
		return
	}

	fast := head
	slow := head
    size := 0
    var lastElementOfFirstHalf *ListNode
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
