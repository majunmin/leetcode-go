package leetcode_1334

import "testing"

func Test_findTheCity(t *testing.T) {
	type args struct {
		n                 int
		edges             [][]int
		distanceThreshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n:                 4,
				edges:             [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
				distanceThreshold: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheCity(tt.args.n, tt.args.edges, tt.args.distanceThreshold); got != tt.want {
				t.Errorf("findTheCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
