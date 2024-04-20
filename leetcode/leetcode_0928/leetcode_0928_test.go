package leetcode_0928

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
					{1, 1, 0},
					{1, 1, 1},
					{0, 1, 1},
				},
				initial: []int{0, 1},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				graph: [][]int{
					{1, 0, 0, 0, 0, 0, 0, 0, 1},
					{0, 1, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 1, 1, 0, 1, 0, 0, 0},
					{0, 1, 1, 1, 1, 0, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 0, 1, 1, 0, 0, 0},
					{0, 0, 0, 1, 0, 0, 1, 1, 0},
					{0, 0, 0, 0, 0, 0, 1, 1, 1},
					{1, 0, 0, 0, 0, 0, 0, 1, 1},
				},
				initial: []int{3, 7},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				graph: [][]int{
					{1, 0, 0, 1},
					{0, 1, 1, 0},
					{0, 1, 1, 0},
					{1, 0, 0, 1},
				},
				initial: []int{3, 1},
			},
			want: 3,
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
