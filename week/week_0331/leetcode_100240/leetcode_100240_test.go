package leetcode_100240

import "testing"

func Test_minimumDistance(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				[][]int{
					{3, 10},
					{5, 15},
					{10, 2},
					{4, 4},
				},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				[][]int{
					{3, 2},
					{3, 9},
					{7, 10},
					{4, 4},
					{8, 10},
					{2, 7},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDistance(tt.args.points); got != tt.want {
				t.Errorf("minimumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
