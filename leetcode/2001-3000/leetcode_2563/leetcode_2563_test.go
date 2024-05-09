package leetcode_2563

import "testing"

func Test_countFairPairs(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				nums:  []int{0, 1, 4, 4, 5, 7},
				lower: 3,
				upper: 6,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums:  []int{-5, -7, -5, -7, -5},
				lower: -12,
				upper: -12,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFairPairs(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countFairPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
