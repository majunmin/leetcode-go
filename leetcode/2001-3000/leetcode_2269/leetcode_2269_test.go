package leetcode_2269

import "testing"

func Test_divisorSubstrings(t *testing.T) {
	type args struct {
		num int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				num: 240,
				k:   2,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				num: 30003,
				k:   3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorSubstrings(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("divisorSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
