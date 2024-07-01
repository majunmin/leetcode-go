package leetcode_2065

import "testing"

func Test_maximalPathQuality(t *testing.T) {
	type args struct {
		values  []int
		edges   [][]int
		maxTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				values: []int{0, 32, 10, 43},
				edges: [][]int{
					{0, 1, 10},
					{1, 2, 15},
					{0, 3, 10},
				},
				maxTime: 49,
			},
			want: 75,
		},
		{
			name: "",
			args: args{
				values: []int{5, 10, 15, 20},
				edges: [][]int{
					{0, 1, 10},
					{1, 2, 15},
					{0, 3, 10},
				},
				maxTime: 30,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalPathQuality(tt.args.values, tt.args.edges, tt.args.maxTime); got != tt.want {
				t.Errorf("maximalPathQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
