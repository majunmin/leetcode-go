package leetcode_0898

import "testing"

func Test_subarrayBitwiseORs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{arr: []int{1, 2, 4}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarrayBitwiseORs(tt.args.arr); got != tt.want {
				t.Errorf("subarrayBitwiseORs() = %v, want %v", got, tt.want)
			}
		})
	}
}
