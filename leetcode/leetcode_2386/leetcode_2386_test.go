package leetcode_2386

import "testing"

func Test_kSum(t *testing.T) {
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
			name: "case0",
			args: args{
				nums: []int{2, 4, -2},
				k:    4,
			},
			want: 2,
		},
		{
			name: "case1",
			args: args{
				nums: []int{1, -2, 3, 4, -10, 12},
				k:    16,
			},
			want: 10,
		},
		{
			name: "case3",
			args: args{
				nums: []int{492634335, 899178915, 899178915},
				k:    2,
			},
			want: 1391813250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("kSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
