package leetcode_2760

import "testing"

func Test_longestAlternatingSubarray(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{nums: []int{3, 2, 5, 4}, threshold: 5},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestAlternatingSubarray(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("longestAlternatingSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
