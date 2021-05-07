package Solution

import (
	"math"
	"tree"
)

func maxPartialPath(root *tree.TreeNode) float64 {
	if root == nil {
		return 0
	}

	maxChildPath := math.Max(
		float64(maxPartialPath(root.Left)),
		float64(maxPartialPath(root.Right)),
	)

	return math.Max(float64(root.Val), float64(root.Val)+maxChildPath)
}

func maxPathSumSlow(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	maxRootPathSum := float64(root.Val)

	maxLeftPartialPathSum := maxPartialPath(root.Left)
	maxRightPartialPathSum := maxPartialPath(root.Right)
	maxRootPathSum = math.Max(maxRootPathSum, maxRootPathSum+maxLeftPartialPathSum)
	maxRootPathSum = math.Max(maxRootPathSum, maxRootPathSum+maxRightPartialPathSum)

	if root.Left != nil {
		maxRootPathSum = math.Max(maxRootPathSum, float64(maxPathSum(root.Left)))
	}

	if root.Right != nil {
		maxRootPathSum = math.Max(maxRootPathSum, float64(maxPathSum(root.Right)))
	}

	return int(maxRootPathSum)
}

func maxPathSumUsingRefHelper(root *tree.TreeNode, maxSum *float64) float64 {
	if root == nil {
		return 0
	}
	maxLeftPath := maxPathSumUsingRefHelper(root.Left, maxSum)
	maxRightPath := maxPathSumUsingRefHelper(root.Right, maxSum)
	maxChildPath := math.Max(maxLeftPath, maxRightPath)
	rootVal := float64(root.Val)

	*maxSum = math.Max(rootVal, *maxSum)
	*maxSum = math.Max(rootVal+maxLeftPath, *maxSum)
	*maxSum = math.Max(rootVal+maxRightPath, *maxSum)
	*maxSum = math.Max(rootVal+maxRightPath+maxLeftPath, *maxSum)

	return math.Max(rootVal+maxChildPath, rootVal)
}

func maxPathSumUsingRef(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := float64(root.Val)
	maxPathSumUsingRefHelper(root, &maxSum)

	return int(maxSum)
}

func calculatePathValues(root *tree.TreeNode) (partialPath float64, fullPath int) {
	if root == nil {
		return 0, 0
	}

	maxLeftPartialPath, maxLeftFullPath := 0.0, root.Val
	if root.Left != nil {
		maxLeftPartialPath, maxLeftFullPath = calculatePathValues(root.Left)
	}

	maxRightPartialPath, maxRightFullPath := 0.0, root.Val
	if root.Right != nil {
		maxRightPartialPath, maxRightFullPath = calculatePathValues(root.Right)
	}

	rootVal := float64(root.Val)
	maxChildPartialPath := math.Max(maxLeftPartialPath, maxRightPartialPath)
	maxPartialPath := math.Max(rootVal, rootVal+maxChildPartialPath)

	maxFullPath := rootVal
	maxFullPath = math.Max(maxFullPath, float64(maxLeftFullPath))
	maxFullPath = math.Max(maxFullPath, float64(maxRightFullPath))
	maxFullPath = math.Max(rootVal+maxLeftPartialPath, maxFullPath)
	maxFullPath = math.Max(rootVal+maxRightPartialPath, maxFullPath)
	maxFullPath = math.Max(rootVal+maxRightPartialPath+maxLeftPartialPath, maxFullPath)

	return maxPartialPath, int(maxFullPath)
}

func maxPathSum(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	_, maxPath := calculatePathValues(root)

	return int(maxPath)
}
