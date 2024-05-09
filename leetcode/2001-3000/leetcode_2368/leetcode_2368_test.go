package leetcode_2368

import "testing"

func Test_reachableNodes(t *testing.T) {
	type args struct {
		n          int
		edges      [][]int
		restricted []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:          7,
				edges:      [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}},
				restricted: []int{4, 5},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n:          5,
				edges:      [][]int{{0, 1}, {2, 1}, {1, 3}, {1, 4}},
				restricted: []int{2},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n: 10,
				edges: [][]int{
					{4, 1}, {1, 3}, {1, 5}, {0, 5},
					{3, 6}, {8, 4}, {5, 7}, {6, 9},
					{3, 2},
				},
				restricted: []int{2, 7},
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				n: 15,
				edges: [][]int{
					{2, 6},
					{2, 13},
					{5, 2},
					{9, 5},
					{6, 0},
					{10, 13},
					{14, 13},
					{8, 6},
					{4, 13},
					{10, 7},
					{8, 1},
					{7, 3},
					{6, 12},
					{11, 0},
				},
				restricted: []int{5, 7, 9},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachableNodes(tt.args.n, tt.args.edges, tt.args.restricted); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
