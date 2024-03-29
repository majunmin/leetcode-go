package leetcode_2730

import (
	"testing"
)

func Test_longestSemiRepetitiveSubstring(t *testing.T) {
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
			args: args{"24489929009"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSemiRepetitiveSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestSemiRepetitiveSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
