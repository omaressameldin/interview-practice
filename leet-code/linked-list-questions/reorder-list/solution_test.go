package Solution

import (
	"list_node"
	"reflect"
	"testing"
)

type testTable []struct {
	input    *list_node.ListNode
	expected *list_node.ListNode
}

func creatTestTable() testTable {
	return testTable{
		{
			list_node.CreateFromSlice([]int{1, 2, 3, 4}),
			list_node.CreateFromSlice([]int{1, 4, 2, 3}),
		},
		{
			list_node.CreateFromSlice([]int{1, 2, 3, 4, 5}),
			list_node.CreateFromSlice([]int{1, 5, 2, 4, 3}),
		},
		{
			list_node.CreateFromSlice([]int{1}),
			list_node.CreateFromSlice([]int{1}),
		},
		{
			list_node.CreateFromSlice([]int{1, 2}),
			list_node.CreateFromSlice([]int{1, 2}),
		},
		{
			list_node.CreateFromSlice([]int{}),
			list_node.CreateFromSlice([]int{}),
		},
	}
}
func TestReorderListWithStack(t *testing.T) {
	for _, row := range creatTestTable() {
		ReorderListWithStack(row.input)
		if !reflect.DeepEqual(row.input, row.expected) {
			t.Errorf("incorrect output, got: %s, want: %s.", row.input, row.expected)
		}
	}
}

func TestReorderListWithoutSpace(t *testing.T) {
	for _, row := range creatTestTable() {
		ReorderListWithoutSpace(row.input)
		if !reflect.DeepEqual(row.input, row.expected) {
			t.Errorf("incorrect output, got: %s, want: %s.", row.input, row.expected)
		}
	}
}
