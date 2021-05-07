package Solution

import "tree"

func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
