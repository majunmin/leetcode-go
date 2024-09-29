package leetcode_2334

import "testing"

func Test_validSubarraySize(t *testing.T) {
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
			args: args{
				nums:      []int{1, 3, 4, 3, 1},
				threshold: 6,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validSubarraySize(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("validSubarraySize() = %v, want %v", got, tt.want)
			}
		})
	}
}
