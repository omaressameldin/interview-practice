package tree

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const nullMarker = "#"
const delimiter = ","

func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	answer := make([]string, 0)
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		isNextLevelNil := true
		for levelSize := queue.Len(); levelSize > 0; levelSize-- {
			elem := queue.Front()
			queue.Remove(elem)
			curNode := elem.Value.(*TreeNode)
			if curNode == nil {
				answer = append(answer, nullMarker)
				continue
			}
			answer = append(answer, fmt.Sprintf("%d", curNode.Val))
			isNextLevelNil = isNextLevelNil && curNode.Left == nil && curNode.Right == nil
			queue.PushBack(curNode.Left)
			queue.PushBack(curNode.Right)
		}
		if isNextLevelNil {
			break
		}
	}

	return strings.Join(answer, delimiter)
}

func Deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	treeSlice := strings.Split(data, delimiter)

	rootVal, err := strconv.Atoi(treeSlice[0])
	if err != nil {
		panic(err)
	}
	root := &TreeNode{
		Val: rootVal,
	}
	queue := list.New()
	queue.PushFront(root)
	for index := 1; index < len(treeSlice); index += 2 {
		elem := queue.Front()
		if elem == nil {
			break
		}
		queue.Remove(elem)
		parent := elem.Value.(*TreeNode)
		if treeSlice[index] != nullMarker {
			curVal, err := strconv.Atoi(treeSlice[index])
			if err != nil {
				panic(err)
			}
			parent.Left = &TreeNode{
				Val: curVal,
			}
			queue.PushBack(parent.Left)
		}
		if treeSlice[index+1] != nullMarker {
			curVal, err := strconv.Atoi(treeSlice[index+1])
			if err != nil {
				panic(err)
			}
			parent.Right = &TreeNode{
				Val: curVal,
			}
			queue.PushBack(parent.Right)
		}
	}

	return root
}

func (tree *TreeNode) String() string {
	return Serialize(tree)
}
