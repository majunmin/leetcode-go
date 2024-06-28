package leetcode_100279

import (
	"reflect"
	"testing"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		n         int
		edges     [][]int
		disappear []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1, 2},
					{0, 2, 4},
					{1, 2, 1},
				},
				disappear: []int{1, 3, 5},
			},
			want: []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dijstraSolution(tt.args.n, tt.args.edges, tt.args.disappear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
