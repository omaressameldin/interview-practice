package Solution

import (
	"testing"
	"tree"
)

type testTable []struct {
	input    *tree.TreeNode
	expected int
}

func createTestTable() testTable {
	return testTable{
		{
			input:    nil,
			expected: 0,
		},
		{
			input:    &tree.TreeNode{Val: -3},
			expected: -3,
		},
		{
			input: &tree.TreeNode{
				Val:  2,
				Left: &tree.TreeNode{Val: -1},
			},
			expected: 2,
		},
		{
			input: &tree.TreeNode{
				Val:  -2,
				Left: &tree.TreeNode{Val: -1},
			},
			expected: -1,
		},
		{
			input:    &tree.TreeNode{Val: 5},
			expected: 5,
		},
		{
			input: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
			expected: 8,
		},
		{
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 1,
					Left: &tree.TreeNode{
						Val:  4,
						Left: &tree.TreeNode{Val: -2},
					},
				},
			},
			expected: 10,
		},
		{
			input: &tree.TreeNode{
				Val:   1,
				Left:  &tree.TreeNode{Val: -2},
				Right: &tree.TreeNode{Val: 3},
			},
			expected: 4,
		},
		{
			input: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val:   1,
					Left:  &tree.TreeNode{Val: 2},
					Right: &tree.TreeNode{Val: 1},
				},
				Right: &tree.TreeNode{
					Val:   2,
					Right: &tree.TreeNode{Val: 3},
					Left:  &tree.TreeNode{Val: 4},
				},
			},
			expected: 14,
		},
		{
			input: &tree.TreeNode{
				Val:  -10,
				Left: &tree.TreeNode{Val: 9},
				Right: &tree.TreeNode{
					Val:   20,
					Right: &tree.TreeNode{Val: 15},
					Left:  &tree.TreeNode{Val: 7},
				},
			},
			expected: 42,
		},
	}
}

func TestMaxPathSum(t *testing.T) {
	for _, row := range createTestTable() {
		outputSlow := maxPathSum(row.input)
		if outputSlow != row.expected {
			t.Errorf("Wrong outputSlow, expected: [%d], got: [%d]", row.expected, outputSlow)
		}

		outputForRef := maxPathSumUsingRef(row.input)
		if outputForRef != row.expected {
			t.Errorf("Wrong outputForRef, expected: [%d], got: [%d]", row.expected, outputForRef)
		}

		output := maxPathSum(row.input)
		if output != row.expected {
			t.Errorf("Wrong output, expected: [%d], got: [%d]", row.expected, output)
		}
	}
}
