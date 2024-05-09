package leetcode_1766

import (
	"reflect"
	"testing"
)

func Test_getCoprimes(t *testing.T) {
	type args struct {
		nums  []int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{2, 3, 3, 2},
				edges: [][]int{
					{0, 1}, {1, 2}, {1, 3},
				},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				nums: []int{5, 6, 10, 2, 3, 6, 15},
				edges: [][]int{
					{0, 1},
					{0, 2},
					{1, 3},
					{1, 4},
					{2, 5},
					{2, 6},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCoprimes(tt.args.nums, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
