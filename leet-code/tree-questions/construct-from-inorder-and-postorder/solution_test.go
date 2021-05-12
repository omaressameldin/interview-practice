package Solution

import (
	"reflect"
	"testing"
	"tree"
)

type testTable []struct {
	preorder []int
	inorder  []int
	expected *tree.TreeNode
}

func createTestTable() testTable {
	return testTable{
		{
			preorder: []int{},
			inorder:  []int{},
			expected: nil,
		},
		{
			preorder: []int{15},
			inorder:  []int{15},
			expected: &tree.TreeNode{Val: 15},
		},
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: &tree.TreeNode{
				Val:  3,
				Left: &tree.TreeNode{Val: 9},
				Right: &tree.TreeNode{
					Val:   20,
					Left:  &tree.TreeNode{Val: 15},
					Right: &tree.TreeNode{Val: 7},
				},
			},
		},
	}
}

func TestBuildTree(t *testing.T) {
	for _, row := range createTestTable() {
		output := buildTree(row.preorder, row.inorder)
		if !reflect.DeepEqual(output, row.expected) {
			t.Errorf("Wrong output, expected: [%v], got: [%v]", row.expected, output)
		}
	}
}
