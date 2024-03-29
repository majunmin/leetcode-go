package leetcode_100245

import "testing"

func Test_maximumLengthSubstring(t *testing.T) {
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
			args: args{"aaaaaa"},
			want: 2,
		},
		{
			name: "case2",
			args: args{"bcbbbcba"},
			want: 4,
		},
		{
			name: "case3",
			args: args{"aadaad"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLengthSubstring(tt.args.s); got != tt.want {
				t.Errorf("maximumLengthSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
