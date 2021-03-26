package tree

import (
	"reflect"
	"testing"
)

type SerializeTestTable []struct {
	deserialized *TreeNode
	serialized   string
}

func createSerializeTestTable() SerializeTestTable {
	return SerializeTestTable{
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
	for _, row := range createSerializeTestTable() {
		result := Serialize(row.deserialized)
		if !reflect.DeepEqual(row.serialized, result) {
			t.Errorf("incorrect output, got: [%s], want: [%s].", result, row.serialized)
		}
	}
}

func TestDeserialize(t *testing.T) {
	for _, row := range createSerializeTestTable() {
		result := Deserialize(row.serialized)
		if !reflect.DeepEqual(row.deserialized, result) {
			t.Errorf("incorrect output, got: [%s], want: [%s].", result, row.deserialized)
		}
	}
}

func TestMaxDepth(t *testing.T) {
	for _, row := range []struct {
		tree          *TreeNode
		expectedDepth int
	}{
		{
			tree:          nil,
			expectedDepth: 0,
		},
		{
			tree:          &TreeNode{Val: 1},
			expectedDepth: 1,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 4},
			},
			expectedDepth: 3,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}},
				Left:  &TreeNode{Val: 4},
			},
			expectedDepth: 3,
		},
	} {
		result := TreeDepth(row.tree)
		if row.expectedDepth != result {
			t.Errorf("incorrect depth, got: %d, want: %d.", result, row.expectedDepth)
		}
	}
}
