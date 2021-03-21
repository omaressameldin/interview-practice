package Solution

import (
	"list_node"
	"reflect"
	"testing"
)

type testTable []struct {
	list1    *list_node.ListNode
	list2    *list_node.ListNode
	expected *list_node.ListNode
}

func creatTestTable() testTable {
	return testTable{
		{
			list_node.CreateFromSlice([]int{1, 2, 3}),
			list_node.CreateFromSlice([]int{4, 5, 6}),
			list_node.CreateFromSlice([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			list_node.CreateFromSlice([]int{1, 2, 3, 4}),
			list_node.CreateFromSlice([]int{2, 5, 7, 8}),
			list_node.CreateFromSlice([]int{1, 2, 2, 3, 4, 5, 7, 8}),
		},
		{
			list_node.CreateFromSlice([]int{1, 2, 3, 4}),
			list_node.CreateFromSlice([]int{1, 5, 7, 8}),
			list_node.CreateFromSlice([]int{1, 1, 2, 3, 4, 5, 7, 8}),
		},
		{
			list_node.CreateFromSlice([]int{1, 9}),
			list_node.CreateFromSlice([]int{1, 5, 7, 8}),
			list_node.CreateFromSlice([]int{1, 1, 5, 7, 8, 9}),
		},
		{
			list_node.CreateFromSlice([]int{}),
			list_node.CreateFromSlice([]int{1, 5, 7, 8}),
			list_node.CreateFromSlice([]int{1, 5, 7, 8}),
		},
		{
			list_node.CreateFromSlice([]int{}),
			list_node.CreateFromSlice([]int{}),
			list_node.CreateFromSlice([]int{}),
		},
	}
}
func TestMergeTwoSortedListsInPlace(t *testing.T) {
	for _, row := range creatTestTable() {
		merged := MergeTwoSortedListsInPlace(row.list1, row.list2)
		if !reflect.DeepEqual(row.expected, merged) {
			t.Errorf("incorrect output, got: %s, want: %s.", merged, row.expected)
		}
	}
}

func TestMergeTwoSortedListsWithExtraSpace(t *testing.T) {
	for _, row := range creatTestTable() {
		merged := MergeTwoSortedListsWithExtraSpace(row.list1, row.list2)
		if !reflect.DeepEqual(row.expected, merged) {
			t.Errorf("incorrect output, got: %s, want: %s.", merged, row.expected)
		}
	}
}

// func TestReorderListWithoutSpace(t *testing.T) {
// 	for _, row := range creatTestTable() {
// 		ReorderListWithoutSpace(row.input)
// 		if !reflect.DeepEqual(row.input, row.expected) {
// 			t.Errorf("incorrect output, got: %s, want: %s.", row.input, row.expected)
// 		}
// 	}
// }
