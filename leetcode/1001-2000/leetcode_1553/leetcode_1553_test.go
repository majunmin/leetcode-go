package leetcode_1553

import "testing"

func Test_minDays2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{10},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays2(tt.args.n); got != tt.want {
				t.Errorf("minDays2() = %v, want %v", got, tt.want)
			}
		})
	}
}
