package leetcode_0679

import "testing"

func Test_judgePoint24(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{[]int{4, 1, 8, 7}},
			want: true,
		},
		{
			name: "",
			args: args{[]int{1, 2, 1, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.args.cards); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}
