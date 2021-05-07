package Solution

import (
	"testing"
	"tree"
)

type TestTable []struct {
	input    *tree.TreeNode
	expected bool
}

func createTestTable() TestTable {
	return TestTable{
		{
			input:    nil,
			expected: true,
		},
		{
			input: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 6},
			},
			expected: true,
		},
		{
			input: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 5},
			},
			expected: false,
		},
		{
			input: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 7},
				Right: &tree.TreeNode{Val: 6},
			},
			expected: false,
		},
		{
			input: &tree.TreeNode{
				Val:  5,
				Left: &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{
					Val:  6,
					Left: &tree.TreeNode{Val: 4},
				},
			},
			expected: false,
		},
	}
}

func TestIsValidBST(t *testing.T) {
	for _, row := range createTestTable() {
		output := isValidBST(row.input)
		if row.expected != output {
			t.Errorf("wrong output expected: [%t], got: [%t]", row.expected, output)
		}
	}
}
