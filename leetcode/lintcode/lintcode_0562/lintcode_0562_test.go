package lintcode_0562

import "testing"

func TestBackPackIV(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{2, 3, 6, 7},
				target: 7,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackPackIV(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BackPackIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
