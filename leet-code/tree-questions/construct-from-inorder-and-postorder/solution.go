package Solution

import (
	"container/list"
	"tree"
)

func extractNodeFromElement(element *list.Element) *tree.TreeNode {
	if element == nil {
		return nil
	}

	nodeInStack, ok := element.Value.(*tree.TreeNode)
	if !ok {
		panic("Something went wrong")
	}

	return nodeInStack
}

func removeFinishedNodes(inorder []int, nodesStack *list.List) []int {
	var prevNode *tree.TreeNode
	for {
		stackPeak := nodesStack.Front()
		nodeInStack := extractNodeFromElement(stackPeak)
		if nodeInStack == nil || nodeInStack.Val != inorder[0] {
			break
		}

		prevNode = nodeInStack
		inorder = inorder[1:]
		nodesStack.Remove(stackPeak)
	}

	if prevNode != nil {
		nodesStack.PushFront(prevNode)
	}

	return inorder
}

func buildTreeHelper(preorder []int, inorder []int, nodesStack *list.List) {
	if len(preorder) == 0 {
		return
	}
	remaining := len(inorder)
	inorder = removeFinishedNodes(inorder, nodesStack)
	noneRemoved := remaining == len(inorder)

	newNode := &tree.TreeNode{Val: preorder[0]}
	stackPeak := nodesStack.Front()
	nodeInStack := extractNodeFromElement(stackPeak)

	if noneRemoved {
		nodeInStack.Left = newNode
	} else {
		nodeInStack.Right = newNode
		nodesStack.Remove(stackPeak)
	}
	nodesStack.PushFront(newNode)

	buildTreeHelper(preorder[1:], inorder, nodesStack)
}

func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	root := &tree.TreeNode{Val: preorder[0]}
	nodesStack := list.New()
	nodesStack.PushFront(root)

	buildTreeHelper(preorder[1:], inorder, nodesStack)

	return root
}
