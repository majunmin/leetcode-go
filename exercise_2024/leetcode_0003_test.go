package exercise_2024

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "case2",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "case3",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "case4",
			args: args{"abc"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
