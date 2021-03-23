package Solution

import (
	"math"
	"tree"
)

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(
		float64(maxDepth(root.Left)),
		float64(maxDepth(root.Right)),
	)) + 1
}
