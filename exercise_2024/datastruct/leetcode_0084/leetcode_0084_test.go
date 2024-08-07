package leetcode_0084

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{[]int{2, 1, 5, 6, 2, 3}},
			want: 10,
		},
		{
			name: "case2",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "case2",
			args: args{[]int{2, 4}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
