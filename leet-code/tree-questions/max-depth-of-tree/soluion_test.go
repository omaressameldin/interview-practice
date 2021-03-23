package Solution

import (
	"testing"
	"tree"
)

type TestTable []struct {
	input         *tree.TreeNode
	expectedDepth int
}

func createTestTable() TestTable {
	return TestTable{
		{
			input:         nil,
			expectedDepth: 0,
		},
		{
			input: &tree.TreeNode{
				Val: 5,
			},
			expectedDepth: 1,
		},
		{
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 3,
					Right: &tree.TreeNode{
						Val: 2,
					},
				},
				Right: &tree.TreeNode{
					Val: 4,
					Right: &tree.TreeNode{
						Val: 4,
					},
				},
			},
			expectedDepth: 3,
		},
		{
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 3,
					Right: &tree.TreeNode{
						Val: 2,
					},
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
			expectedDepth: 3,
		},
		{
			input: &tree.TreeNode{
				Val: 5,
				Right: &tree.TreeNode{
					Val: 3,
					Left: &tree.TreeNode{
						Val: 2,
					},
				},
				Left: &tree.TreeNode{
					Val: 4,
				},
			},
			expectedDepth: 3,
		},
	}
}

func TestMaxDepth(t *testing.T) {
	for _, row := range createTestTable() {
		output := maxDepth(row.input)
		if output != row.expectedDepth {
			t.Errorf("Wrong depth, expected: %d, got: %d", row.expectedDepth, output)
		}
	}
}
