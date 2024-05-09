package leetcode_2411

import (
	"reflect"
	"testing"
)

func Test_smallestSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{[]int{1, 0, 2, 1, 3}},
			want: []int{3, 3, 2, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubarrays(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
