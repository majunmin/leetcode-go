package leetcode_2581

import "testing"

func Test_rootCount(t *testing.T) {
	type args struct {
		edges   [][]int
		guesses [][]int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				edges: [][]int{
					{0, 1}, {1, 2}, {1, 3}, {4, 2},
				},
				guesses: [][]int{
					{1, 3}, {0, 1}, {1, 0}, {2, 4},
				},
				k: 4,
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				edges: [][]int{
					{0, 1}, {2, 0}, {0, 3}, {4, 2}, {3, 5}, {6, 0}, {1, 7}, {2, 8}, {2, 9},
					{4, 10}, {9, 11}, {3, 12}, {13, 8}, {14, 9}, {15, 9}, {10, 16},
				},
				guesses: [][]int{
					{8, 2}, {12, 3}, {0, 1}, {16, 10},
				},
				k: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rootCount(tt.args.edges, tt.args.guesses, tt.args.k); got != tt.want {
				t.Errorf("rootCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
