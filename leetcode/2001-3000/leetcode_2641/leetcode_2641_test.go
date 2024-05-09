package leetcode_2641

import (
	"reflect"
	"testing"

	. "github.com/majunmin/leetcode-go/common"
)

func Test_replaceValueInTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   10,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  9,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceValueInTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceValueInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
