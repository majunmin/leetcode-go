package leetcode_100280

import "testing"

func Test_minRectanglesToCoverPoints(t *testing.T) {
	type args struct {
		points [][]int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				points: [][]int{
					{0, 0},
					{1, 1},
					{2, 2},
					{3, 3},
					{4, 4},
					{5, 5},
					{6, 6},
				},
				w: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRectanglesToCoverPoints(tt.args.points, tt.args.w); got != tt.want {
				t.Errorf("minRectanglesToCoverPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
