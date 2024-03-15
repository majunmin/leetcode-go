package leetcode_1976

import "testing"

func Test_countPaths(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 7,
				roads: [][]int{
					{0, 6, 7},
					{0, 1, 2},
					{1, 2, 3},
					{1, 3, 3},
					{6, 3, 3},
					{3, 5, 1},
					{6, 5, 1},
					{2, 5, 1},
					{0, 4, 5},
					{4, 6, 2},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPaths2(tt.args.n, tt.args.roads); got != tt.want {
				t.Errorf("countPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
