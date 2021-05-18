package Solution

import "sort"

func canBeEqualBySort(arrA []int, arrB []int) bool {
	sort.Ints(arrA)
	sort.Ints(arrB)

	for idx := 0; idx < len(arrA); idx += 1 {
		if arrA[idx] != arrB[idx] {
			return false
		}
	}

	return true
}

func canBeEqualByMap(arrA []int, arrB []int) bool {
	elements := make(map[int]int, len(arrA))
	for _, elem := range arrA {
		elements[elem] += 1
	}

	for _, elem := range arrB {
		if elements[elem] == 0 {
			return false
		}
		elements[elem] -= 1
	}

	return true
}
