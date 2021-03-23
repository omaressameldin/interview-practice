package Solution

import "tree"

func IsSameTree(t1, t2 *tree.TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}
	if t2 == nil {
		return false
	}

	return t1.Val == t2.Val &&
		IsSameTree(t1.Left, t2.Left) &&
		IsSameTree(t1.Right, t2.Right)
}
