package Solution

import (
	"reflect"
	"testing"
	"tree"
)

type testTable []struct {
	input    *tree.TreeNode
	expected *tree.TreeNode
}

func createTestTable() testTable {
	return testTable{
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    &tree.TreeNode{Val: 5},
			expected: &tree.TreeNode{Val: 5},
		},
		{
			input: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 2},
				Right: &tree.TreeNode{Val: 1},
			},
			expected: &tree.TreeNode{
				Val:   5,
				Left:  &tree.TreeNode{Val: 1},
				Right: &tree.TreeNode{Val: 2},
			},
		},
		{
			input: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val:   2,
					Left:  &tree.TreeNode{Val: 1},
					Right: &tree.TreeNode{Val: 3},
				},
				Right: &tree.TreeNode{
					Val:   7,
					Left:  &tree.TreeNode{Val: 6},
					Right: &tree.TreeNode{Val: 9},
				},
			},
			expected: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val:   7,
					Left:  &tree.TreeNode{Val: 9},
					Right: &tree.TreeNode{Val: 6},
				},
				Right: &tree.TreeNode{
					Val:   2,
					Left:  &tree.TreeNode{Val: 3},
					Right: &tree.TreeNode{Val: 1},
				},
			},
		},
	}
}

func TestInvertTree(t *testing.T) {
	for _, row := range createTestTable() {
		output := invertTree(row.input)
		if !reflect.DeepEqual(output, row.expected) {
			t.Errorf("Wrong output, expected: [%v], got: [%v]", row.expected, output)
		}
	}
}
