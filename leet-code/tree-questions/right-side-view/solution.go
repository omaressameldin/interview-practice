package Solution

import (
	"container/list"
	"tree"
)

func removeNodeFromQueue(queue *list.List) *tree.TreeNode {
	element := queue.Front()
	queue.Remove(element)
	curNode, ok := element.Value.(*tree.TreeNode)
	if !ok {
		panic("Something went wrong while typecasting")
	}

	return curNode
}

func addNextLevelNodesToQueue(queue *list.List, node *tree.TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		queue.PushBack(node.Left)
	}
	if node.Right != nil {
		queue.PushBack(node.Right)
	}
}

func rightSideView(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	visibleNodes := make([]int, 0, 0)
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		for remainingInLevel := queue.Len(); remainingInLevel > 0; remainingInLevel -= 1 {
			curNode := removeNodeFromQueue(queue)
			addNextLevelNodesToQueue(queue, curNode)
			if remainingInLevel == 1 {
				visibleNodes = append(visibleNodes, curNode.Val)
			}
		}
	}

	return visibleNodes
}
