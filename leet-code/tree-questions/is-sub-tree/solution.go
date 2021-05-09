package Solution

import "tree"

func isSameTree(t1, t2 *tree.TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return false
	}

	return t1.Val == t2.Val &&
		isSameTree(t1.Left, t2.Left) &&
		isSameTree(t1.Right, t2.Right)
}

func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
