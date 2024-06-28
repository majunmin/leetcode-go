package leetcode_2741

import "testing"

func Test_specialPerm(t *testing.T) {
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
			args: args{[]int{2, 3, 6}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialPerm(tt.args.nums); got != tt.want {
				t.Errorf("specialPerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
