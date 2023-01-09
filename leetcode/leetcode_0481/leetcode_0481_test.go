package leetcode_0481

import "testing"

func Test_magicalString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: struct{ n int }{n: 3}, want: 1},
		{name: "case2", args: struct{ n int }{n: 6}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicalString(tt.args.n); got != tt.want {
				t.Errorf("magicalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
