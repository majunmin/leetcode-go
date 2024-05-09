package leetcode_2680

import "testing"

func Test_maximumOr(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				nums: []int{12, 9},
				k:    1,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOr(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
