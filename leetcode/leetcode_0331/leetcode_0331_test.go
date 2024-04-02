package leetcode_0331

import "testing"

func Test_stackSolution(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{"9,3,4,#,#,1,#,#,2,#,6,#,#"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stackSolution(tt.args.preorder); got != tt.want {
				t.Errorf("stackSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
