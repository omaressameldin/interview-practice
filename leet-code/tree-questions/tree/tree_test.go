package tree

import (
	"reflect"
	"testing"
)

type testTable []struct {
	deserialized *TreeNode
	serialized   string
}

func createTestTable() testTable {
	return testTable{
		{
			deserialized: nil,
			serialized:   "",
		},
		{
			deserialized: &TreeNode{Val: 1},
			serialized:   "1",
		},
		{
			deserialized: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			serialized: "1,2,3",
		},
		{
			deserialized: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			serialized: "1,2,3,#,#,4,#",
		},
		{
			deserialized: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			serialized: "1,2,#,3,#",
		},
		{
			serialized: "1,2,3,#,#,4,5,6,7,#,#",
			deserialized: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
}

func TestSerialize(t *testing.T) {
	for _, row := range createTestTable() {
		result := Serialize(row.deserialized)
		if !reflect.DeepEqual(row.serialized, result) {
			t.Errorf("incorrect output, got: [%s], want: [%s].", result, row.serialized)
		}
	}
}

func TestDeserialize(t *testing.T) {
	for _, row := range createTestTable() {
		result := Deserialize(row.serialized)
		if !reflect.DeepEqual(row.deserialized, result) {
			t.Errorf("incorrect output, got: [%s], want: [%s].", result, row.deserialized)
		}
	}
}
