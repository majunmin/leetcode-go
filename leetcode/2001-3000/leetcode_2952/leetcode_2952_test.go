package leetcode_2952

import "testing"

func Test_minimumAddedCoins(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				coins:  []int{1, 4, 10},
				target: 19,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAddedCoins(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("minimumAddedCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
