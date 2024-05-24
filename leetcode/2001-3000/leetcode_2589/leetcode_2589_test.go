package leetcode_2589

import "testing"

func Test_findMinimumTime(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{tasks: [][]int{
				{2, 3, 1},
				{4, 5, 1},
				{1, 5, 2},
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumTime(tt.args.tasks); got != tt.want {
				t.Errorf("findMinimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
