package leetcode_3283

import "testing"

func Test_maxMoves(t *testing.T) {
	type args struct {
		kx        int
		ky        int
		positions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				kx:        1,
				ky:        1,
				positions: [][]int{{0, 0}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMoves(tt.args.kx, tt.args.ky, tt.args.positions); got != tt.want {
				t.Errorf("maxMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
