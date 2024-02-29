package leetcode_0224

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{"1+2+3-4"},
			want: 2,
		},
		{
			name: "case2",
			args: args{"-1-2-3"},
			want: -6,
		},
		{
			name: "case3",
			args: args{"-1-(2-3)"},
			want: 0,
		},
		{
			name: "case4",
			args: args{"(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		},
		{
			name: "case5",
			args: args{"1-( -2)"},
			want: 3,
		},
		{
			name: "case6",
			args: args{"1-( 3-2*4)"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
