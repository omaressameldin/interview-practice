package Solution

import (
	"testing"
	"tree"
)

type TestTable []struct {
	mainTree *tree.TreeNode
	subTree  *tree.TreeNode
	expected bool
}

func createTestTable() TestTable {
	return TestTable{
		{
			mainTree: nil,
			subTree:  nil,
			expected: true,
		},
		{
			mainTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			subTree:  nil,
			expected: true,
		},
		{
			mainTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			subTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			mainTree: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val:   5,
					Left:  &tree.TreeNode{Val: 1},
					Right: &tree.TreeNode{Val: 2},
				},
				Right: &tree.TreeNode{Val: 2},
			},
			subTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			mainTree: &tree.TreeNode{
				Val:  5,
				Left: &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{
					Val:   5,
					Left:  &tree.TreeNode{Val: 1},
					Right: &tree.TreeNode{Val: 2},
				},
			},
			subTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			mainTree: &tree.TreeNode{
				Val:   5,
				Right: &tree.TreeNode{Val: 1},
				Left: &tree.TreeNode{
					Val:   5,
					Left:  &tree.TreeNode{Val: 1},
					Right: &tree.TreeNode{Val: 2},
				},
			},
			subTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			mainTree: &tree.TreeNode{
				Val:   5,
				Right: &tree.TreeNode{Val: 1},
				Left: &tree.TreeNode{
					Val: 5,
					Left: &tree.TreeNode{
						Val:   1,
						Right: &tree.TreeNode{Val: 2},
					},
					Right: &tree.TreeNode{Val: 2},
				},
			},
			subTree: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			mainTree: &tree.TreeNode{
				Val: 3,
				Right: &tree.TreeNode{
					Val:  5,
					Left: &tree.TreeNode{Val: 2},
				},
				Left: &tree.TreeNode{
					Val:  4,
					Left: &tree.TreeNode{Val: 1},
				},
			},
			subTree: &tree.TreeNode{
				Val:   3,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: false,
		},
	}
}

func TestIsSubTree(t *testing.T) {
	for _, row := range createTestTable() {
		output := isSubtree(row.mainTree, row.subTree)
		if output != row.expected {
			t.Errorf("Wrong output, expected: %t, got: %t", row.expected, output)
		}
	}
}
