package list_node

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateFromSlice(items []int) *ListNode {
	if len(items) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  items[0],
		Next: nil,
	}
	cur := head
	for _, item := range items[1:] {
		cur.Next = &ListNode{
			Val:  item,
			Next: nil,
		}
		cur = cur.Next
	}

	return head
}

func (list *ListNode) String() string {
	if list == nil {
		return "NONE"
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d", list.Val))
	cur := list.Next
	for cur != nil {
		b.WriteString(fmt.Sprintf("->%d", cur.Val))
		cur = cur.Next
	}

	return b.String()
}
