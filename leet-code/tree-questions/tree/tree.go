package tree

import (
	"container/list"
	"fmt"
	"math"
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
	for depth := TreeDepth(root); depth > 0; depth-- {
		nextLevelQueue, serializedLevel := serializeTreeLevel(queue)
		queue = nextLevelQueue
		answer = append(answer, serializedLevel...)
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
	root := &TreeNode{Val: rootVal}
	queue := list.New()
	queue.PushFront(root)
	for index := 1; index < len(treeSlice); index += 2 {
		elem := queue.Front()
		queue.Remove(elem)
		parent := elem.Value.(*TreeNode)
		parent.Left = deserializeElement(treeSlice[index])
		if parent.Left != nil {
			queue.PushBack(parent.Left)
		}
		parent.Right = deserializeElement(treeSlice[index+1])
		if parent.Right != nil {
			queue.PushBack(parent.Right)
		}
	}

	return root
}

func (tree *TreeNode) String() string {
	return Serialize(tree)
}

func serializeTreeLevel(queue *list.List) (nextLevelQueue *list.List, serialziedLevel []string) {
	serialziedLevel = make([]string, 0)
	nextLevelQueue = list.New()
	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		curNode := elem.Value.(*TreeNode)
		if curNode == nil {
			serialziedLevel = append(serialziedLevel, nullMarker)
			continue
		}
		serialziedLevel = append(serialziedLevel, fmt.Sprintf("%d", curNode.Val))
		nextLevelQueue.PushBack(curNode.Left)
		nextLevelQueue.PushBack(curNode.Right)
	}

	return nextLevelQueue, serialziedLevel
}

func deserializeElement(element string) *TreeNode {
	if element == nullMarker {
		return nil
	}

	curVal, err := strconv.Atoi(element)
	if err != nil {
		panic(err)
	}

	return &TreeNode{
		Val: curVal,
	}
}

func TreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(
		float64(TreeDepth(root.Left)),
		float64(TreeDepth(root.Right)),
	)) + 1
}
