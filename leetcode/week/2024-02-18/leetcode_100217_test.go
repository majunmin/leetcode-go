package _024_02_18

import (
	"testing"
)

func Test_mostFrequentPrime(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{[][]int{
				{1, 1},
				{9, 9},
				{1, 1},
			}},
			want: 19,
		},
		{
			name: "case2",
			args: args{[][]int{
				{9, 7, 8},
				{4, 6, 5},
				{2, 8, 6},
			}},
			want: 97,
		},
		{
			name: "case3",
			args: args{[][]int{
				{9, 3, 8},
				{4, 2, 5},
				{3, 8, 6},
			}},
			want: 823,
		},
		{
			name: "case3",
			args: args{[][]int{
				{3, 9, 7},
				{9, 7, 4},
				{4, 5, 4},
			}},
			want: 47,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequentPrime(tt.args.mat); got != tt.want {
				t.Errorf("mostFrequentPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
