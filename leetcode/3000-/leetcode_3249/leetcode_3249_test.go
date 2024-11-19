package leetcode_3249

import "testing"

func Test_countGoodNodes(t *testing.T) {
	type args struct {
		edges [][]int
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
					{0, 1},
					{0, 2},
					{1, 3},
					{1, 4},
					{2, 5},
					{2, 6},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodNodes(tt.args.edges); got != tt.want {
				t.Errorf("countGoodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
