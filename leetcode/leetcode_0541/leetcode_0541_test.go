package leetcode_0541

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "abcdefg", k: 2},
			want: "bacdfeg",
		},
		{
			name: "case2",
			args: args{s: "abcd", k: 4},
			want: "dcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(s[12:45])
}
