package leetcode_0394

import "testing"

func Test_decodeString(t *testing.T) {
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
			args: args{"3[a]3[b]"},
			want: "aaabbb",
		},
		{
			name: "case1",
			args: args{""},
			want: "",
		},
		{
			name: "case1",
			args: args{"3[3[3[a]]b]"},
			want: "aaaaaaaaabaaaaaaaaabaaaaaaaaab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
