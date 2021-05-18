package Solution

import (
	"testing"
)

type TestTable []struct {
	brackets string
	expected bool
}

func createTestTable() TestTable {
	return TestTable{
		{
			brackets: "",
			expected: true,
		},
		{
			brackets: "[]",
			expected: true,
		},
		{
			brackets: "[]]",
			expected: false,
		},
		{
			brackets: "[(}]",
			expected: false,
		},
		{
			brackets: "()[]{}",
			expected: true,
		},
		{
			brackets: "([]{[]})",
			expected: true,
		},
	}
}

func TestIsSubTree(t *testing.T) {
	for _, row := range createTestTable() {
		output := isValid(row.brackets)
		if output != row.expected {
			t.Errorf("Wrong output for %s, expected: %t, got: %t", row.brackets, row.expected, output)
		}
	}
}
