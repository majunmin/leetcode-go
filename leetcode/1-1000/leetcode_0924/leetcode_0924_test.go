package leetcode_0924

import "testing"

func Test_minMalwareSpread(t *testing.T) {
	type args struct {
		graph   [][]int
		initial []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				graph: [][]int{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
				initial: []int{0, 2},
			},
			want: 0,
		},
		{
			name: "case1",
			args: args{
				graph: [][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
				initial: []int{0, 1, 2},
			},
			want: 2,
		},
		{
			name: "case1",
			args: args{
				graph: [][]int{
					{1, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 1, 1},
					{0, 0, 1, 1},
				},
				initial: []int{1, 3},
			},
			want: 2,
		},

		{
			name: "case1",
			args: args{
				graph: [][]int{
					{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
					{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
					{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0},
					{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0},
					{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
					{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
				},
				initial: []int{7, 8, 6, 2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMalwareSpread(tt.args.graph, tt.args.initial); got != tt.want {
				t.Errorf("minMalwareSpread() = %v, want %v", got, tt.want)
			}
		})
	}
}
