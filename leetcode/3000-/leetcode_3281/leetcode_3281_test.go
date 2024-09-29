package leetcode_3281

import "testing"

func Test_maxPossibleScore(t *testing.T) {
	type args struct {
		start []int
		d     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				start: []int{3, 0, 6},
				d:     2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPossibleScore(tt.args.start, tt.args.d); got != tt.want {
				t.Errorf("maxPossibleScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
