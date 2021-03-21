package Solution

import (
	"reflect"
	"testing"
)

type testTable []struct {
	input    *ListNode
	expected *ListNode
}

func creatTestTable() testTable {
	return testTable{
		{
			CreateFromSlice([]int{1, 2, 3, 4}),
			CreateFromSlice([]int{1, 4, 2, 3}),
		},
		{
			CreateFromSlice([]int{1, 2, 3, 4, 5}),
			CreateFromSlice([]int{1, 5, 2, 4, 3}),
		},
		{
			CreateFromSlice([]int{1}),
			CreateFromSlice([]int{1}),
		},
		{
			CreateFromSlice([]int{1, 2}),
			CreateFromSlice([]int{1, 2}),
		},
		{
			CreateFromSlice([]int{}),
			CreateFromSlice([]int{}),
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
