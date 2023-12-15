package leetcode_0238

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
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
			args: args{nums: []int{1}},
			want: nil,
		},
		{
			name: "case2",
			args: args{nums: []int{1, 2}},
			want: nil,
		},
		{
			name: "case2",
			args: args{nums: []int{1, 2, 3}},
			want: nil,
		},
		{
			name: "case4",
			args: args{nums: []int{1, 2, 3, 4}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
