package leetcode_100263

import "testing"

func Test_sumOfTheDigitsOfHarshadNumber(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{x: 18},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfTheDigitsOfHarshadNumber(tt.args.x); got != tt.want {
				t.Errorf("sumOfTheDigitsOfHarshadNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
