package leetcode_2917

import "testing"

func Test_findKOr(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{7, 12, 9, 8, 9, 15},
				k:    4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKOr(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
