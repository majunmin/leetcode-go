package leetcode_2653

import (
	"reflect"
	"testing"
)

func Test_getSubarrayBeauty(t *testing.T) {
	type args struct {
		nums []int
		k    int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, -1, -3, -2, 3},
				k:    3,
				x:    2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubarrayBeauty(tt.args.nums, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubarrayBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
