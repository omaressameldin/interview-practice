package Solution

import (
	"reflect"
	"testing"
	"tree"
)

type testTable []struct {
	input    *tree.TreeNode
	expected []int
}

func createTestTable() testTable {
	return testTable{
		{
			input:    nil,
			expected: []int{},
		},
		{
			input:    &tree.TreeNode{Val: 5},
			expected: []int{5},
		},
		{
			input: &tree.TreeNode{
				Val: 8,
				Left: &tree.TreeNode{
					Val: 10,
					Left: &tree.TreeNode{
						Val: 14,
						Left: &tree.TreeNode{
							Val:   6,
							Left:  &tree.TreeNode{Val: 4},
							Right: &tree.TreeNode{Val: 7},
						},
					},
				},
			},
			expected: []int{8, 10, 14, 6, 7},
		},
		{
			input: &tree.TreeNode{
				Val: 8,
				Right: &tree.TreeNode{
					Val: 10,
					Right: &tree.TreeNode{
						Val: 14,
						Right: &tree.TreeNode{
							Val:   6,
							Left:  &tree.TreeNode{Val: 4},
							Right: &tree.TreeNode{Val: 7},
						},
					},
				},
			},
			expected: []int{8, 10, 14, 6, 7},
		},
		{
			input: &tree.TreeNode{
				Val: 8,
				Left: &tree.TreeNode{
					Val:  3,
					Left: &tree.TreeNode{Val: 1},
					Right: &tree.TreeNode{
						Val:   6,
						Left:  &tree.TreeNode{Val: 4},
						Right: &tree.TreeNode{Val: 7},
					},
				},
				Right: &tree.TreeNode{
					Val: 10,
					Right: &tree.TreeNode{
						Val:  14,
						Left: &tree.TreeNode{Val: 13},
					},
				},
			},
			expected: []int{8, 10, 14, 13},
		},
	}
}

func TestRightSideView(t *testing.T) {
	for _, row := range createTestTable() {
		output := rightSideView(row.input)
		if !reflect.DeepEqual(output, row.expected) {
			t.Errorf("Wrong output, expected: %v, got: %v", row.expected, output)
		}
	}
}
