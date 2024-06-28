package leetcode_100281

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{[][]int{
				{9, 5, 7, 3},
				{8, 9, 6, 1},
				{6, 7, 14, 3},
				{2, 5, 3, 1},
			}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.grid); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
