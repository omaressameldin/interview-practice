package Solution

import (
	"reflect"
	"testing"
	"tree"
)

type TestTable []struct {
	input    *tree.TreeNode
	expected [][]int
}

func createTestTable() TestTable {
	return TestTable{
		{
			input:    nil,
			expected: [][]int{},
		},
		{
			input:    &tree.TreeNode{Val: 5},
			expected: [][]int{{5}},
		},
		{
			input: &tree.TreeNode{
				Val:  5,
				Left: &tree.TreeNode{Val: 2},
				Right: &tree.TreeNode{
					Val:  3,
					Left: &tree.TreeNode{Val: 2},
				},
			},
			expected: [][]int{{5}, {2, 3}, {2}},
		},
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	for _, row := range createTestTable() {
		output := levelOrderTraversal(row.input)
		if !reflect.DeepEqual(row.expected, output) {
			t.Errorf("Wrong output, expected: %v, got: %v", row.expected, output)

		}
	}
}
