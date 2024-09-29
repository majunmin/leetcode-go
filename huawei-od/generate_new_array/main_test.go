package main

import "testing"

func Test_countWays(t *testing.T) {
	type args struct {
		M int
		N []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				M: 5,
				N: []int{2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWays2(tt.args.M, tt.args.N); got != tt.want {
				t.Errorf("countWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
