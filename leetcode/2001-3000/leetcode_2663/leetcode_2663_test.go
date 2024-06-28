package leetcode_2663

import "testing"

func Test_smallestBeautifulString(t *testing.T) {
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
			name: "",
			args: args{
				s: "abcz",
				k: 4,
			},
			want: "abda",
		},
		{
			name: "",
			args: args{
				s: "b",
				k: 6,
			},
			want: "c",
		},
		{
			name: "",
			args: args{
				s: "cegaf",
				k: 7,
			},
			want: "cegba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestBeautifulString(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("smallestBeautifulString() = %v, want %v", got, tt.want)
			}
		})
	}
}
