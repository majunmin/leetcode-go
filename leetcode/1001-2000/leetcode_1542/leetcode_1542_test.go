package leetcode_1542

import "testing"

func Test_longestAwesome(t *testing.T) {
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
			args: args{"940884"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestAwesome(tt.args.s); got != tt.want {
				t.Errorf("longestAwesome() = %v, want %v", got, tt.want)
			}
		})
	}
}
