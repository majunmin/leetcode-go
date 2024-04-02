package leetcode_0894

import (
	"reflect"
	"testing"

	"github.com/majunmin/leetcode-go/common"
)

func Test_allPossibleFBT(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*common.TreeNode
	}{
		{
			name: "case1",
			args: args{3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}
