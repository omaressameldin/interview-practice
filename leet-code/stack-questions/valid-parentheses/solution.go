package Solution

import (
	"container/list"
)

var bracketsMap = map[rune]rune{'(': ')', '{': '}', '[': ']'}

func matchClosingBracketWithTop(closingBracket rune, openBracketsStack *list.List) bool {
	stackPeak := openBracketsStack.Front()
	if stackPeak == nil {
		return false
	}

	openBracket, ok := stackPeak.Value.(rune)
	if !ok {
		return false
	}

	openBracketsStack.Remove(stackPeak)

	return bracketsMap[openBracket] == closingBracket
}

func isValid(brackets string) bool {
	openBracketsStack := list.New()
	for _, bracket := range brackets {
		if _, ok := bracketsMap[bracket]; ok {
			openBracketsStack.PushFront(bracket)
		} else if !matchClosingBracketWithTop(bracket, openBracketsStack) {
			return false
		}
	}

	return openBracketsStack.Len() == 0
}
