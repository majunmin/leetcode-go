package leetcode_0316

import "testing"

func Test_removeDuplicateLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{"abcdbc"},
			want: "abcd",
		},
		{
			name: "case2",
			args: args{"bcdabcd"},
			want: "abcd",
		},
		{
			name: "case3",
			args: args{"bcdadcb"},
			want: "adcb",
		},
		{
			name: "case4",
			args: args{"bcabc"},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
