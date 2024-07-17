package leetcode_1617

import "testing"

func Test_minimumMountainRemovals(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{[]int{1, 3, 1}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{9, 8, 1, 7, 6, 5, 4, 3, 2, 1}},
			want: 2,
		},
		{
			name: "",
			args: args{[]int{100, 92, 89, 77, 74, 66, 64, 66, 64}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMountainRemovals(tt.args.nums); got != tt.want {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
