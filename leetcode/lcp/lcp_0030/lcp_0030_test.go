package lcp_0030

import "testing"

func Test_magicTower(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{nums: []int{-1, -1, 10}},
			want: 2,
		},
		{
			name: "case1",
			args: args{nums: []int{-9635, 71923, -37495, 8366, 54303, -86422, -48303, -47858, 98424}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicTower(tt.args.nums); got != tt.want {
				t.Errorf("magicTower() = %v, want %v", got, tt.want)
			}
		})
	}
}
