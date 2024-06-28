package leetcode_100276

import (
	"reflect"
	"testing"
)

func Test_findAnswer(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "case1",
			args: args{
				n: 6,
				edges: [][]int{
					{0, 1, 4},
					{0, 2, 1},
					{1, 3, 2},
					{1, 4, 3},
					{1, 5, 1},
					{2, 3, 1},
					{3, 5, 3},
					{4, 5, 2},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnswer(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}
