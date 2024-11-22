package leetcode_3233

import "testing"

func Test_nonSpecialCount(t *testing.T) {
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				l: 1,
				r: 4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nonSpecialCount(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("nonSpecialCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
