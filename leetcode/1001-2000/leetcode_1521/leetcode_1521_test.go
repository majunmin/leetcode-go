package leetcode_1521

import "testing"

func Test_closestToTarget(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr:    []int{9, 12, 3, 7, 15},
				target: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr:    []int{5, 89, 79, 44, 45, 79},
				target: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestToTarget(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("closestToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
