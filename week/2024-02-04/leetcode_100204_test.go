package _024_02_04

import "testing"

func Test_minimumTimeToInitialState(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{
				word: "abcbabcd",
				k:    2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeToInitialState(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeToInitialState() = %v, want %v", got, tt.want)
			}
		})
	}
}
