package Solution

import (
	"testing"
)

type TestTable []struct {
	arrA     []int
	arrB     []int
	expected bool
}

func createTestTable() TestTable {
	return TestTable{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 4, 3, 2},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 5, 3, 2},
			false,
		},
		{
			[]int{1, 1, 1, 4},
			[]int{1, 1, 2, 4},
			false,
		},
		{
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
			true,
		},
	}
}

func TestIsSubTree(t *testing.T) {
	for _, row := range createTestTable() {
		output := canBeEqualBySort(row.arrA, row.arrB)
		if output != row.expected {
			t.Errorf("Wrong output for %v, %v, expected: %t, got: %t", row.arrA, row.arrB, row.expected, output)
		}
		output = canBeEqualByMap(row.arrA, row.arrB)
		if output != row.expected {
			t.Errorf("Wrong output for %v, %v, expected: %t, got: %t", row.arrA, row.arrB, row.expected, output)
		}
	}
}
