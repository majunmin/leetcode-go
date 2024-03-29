package leetocde_0235

import (
	"reflect"
	"testing"

	. "github.com/majunmin/leetcode-go/common"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val:   3,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   5,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   9,
							Left:  nil,
							Right: nil,
						},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 8},
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				p: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
