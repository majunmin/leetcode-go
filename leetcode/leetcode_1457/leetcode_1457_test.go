package leetcode_1457

import (
	"testing"

	. "github.com/majunmin/leetcode-go/common"
)

var (
	node11 = &TreeNode{Val: 3}
	node12 = &TreeNode{Val: 1}
	node1  = &TreeNode{Val: 3, Left: node11, Right: node12}
	node22 = &TreeNode{Val: 1}
	node2  = &TreeNode{Val: 1, Right: node22}

	node = &TreeNode{Val: 2, Left: node1, Right: node2}
)

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{root: node},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
