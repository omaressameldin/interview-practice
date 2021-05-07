package Solution

import (
	"math"
	"tree"
)

type bstRange struct {
	Min int
	Max int
}

func isBSTInRange(node *tree.TreeNode, treeRange bstRange) bool {
	if node == nil {
		return true
	}

	isNodeInRange := node.Val >= treeRange.Min && node.Val <= treeRange.Max
	return isNodeInRange &&
		isBSTInRange(node.Left, bstRange{treeRange.Min, node.Val - 1}) &&
		isBSTInRange(node.Right, bstRange{node.Val + 1, treeRange.Max})
}

func isValidBST(root *tree.TreeNode) bool {
	return isBSTInRange(root, bstRange{math.MinInt32, math.MaxInt32})
}
