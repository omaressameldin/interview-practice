package Solution

import (
	"testing"
	"tree"
)

type testTable []struct {
	input1   *tree.TreeNode
	input2   *tree.TreeNode
	expected bool
}

func createTestTable() testTable {
	return testTable{
		{
			input1: &tree.TreeNode{
				Val: 5,
			},
			input2: &tree.TreeNode{
				Val: 5,
			},
			expected: true,
		},
		{
			input1:   nil,
			input2:   nil,
			expected: true,
		},
		{
			input1: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 5,
					Right: &tree.TreeNode{
						Val: 6,
						Right: &tree.TreeNode{
							Val: 6,
						},
					},
				},
			},
			input2: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 5,
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
			},
			expected: false,
		},
		{
			input1: &tree.TreeNode{
				Val: 5,
			},
			input2: &tree.TreeNode{
				Val: 5,
			},
			expected: true,
		},
		{
			input1: nil,
			input2: &tree.TreeNode{
				Val: 5,
			},
			expected: false,
		},
		{
			input1: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 5,
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
			},
			input2: &tree.TreeNode{
				Val: 1,
				Right: &tree.TreeNode{
					Val: 5,
					Left: &tree.TreeNode{
						Val: 6,
					},
				},
			},
			expected: false,
		},
	}
}

func TestIsSameTree(t *testing.T) {
	for _, row := range createTestTable() {
		output := IsSameTree(row.input1, row.input2)
		if row.expected != output {
			t.Errorf("Wrong output, expected: %v, got: %v", row.expected, output)
		}
	}
}
