package Solution

import (
	"container/list"
	"tree"
)

func addNextLevelNodesToQueue(node *tree.TreeNode, queue *list.List) {
	if node.Left != nil {
		queue.PushBack(node.Left)
	}
	if node.Right != nil {
		queue.PushBack(node.Right)
	}
}

func getLevelFromQuue(queue *list.List) []int {
	level := make([]int, 0, queue.Len())
	for levelSize := queue.Len(); levelSize > 0; levelSize-- {
		elem := queue.Front()
		node, ok := elem.Value.(*tree.TreeNode)
		if !ok {
			panic("something went wrong trying to extract node")
		}
		level = append(level, node.Val)
		queue.Remove(elem)
		addNextLevelNodesToQueue(node, queue)
	}

	return level
}

func levelOrderTraversal(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := list.New()
	levels := [][]int{}

	queue.PushBack(root)
	for queue.Len() > 0 {
		level := getLevelFromQuue(queue)
		levels = append(levels, level)
	}

	return levels
}
