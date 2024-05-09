package leetcode_0028

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				haystack: "sadfyusad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bmSolution(t *testing.T) {
	idx := boyerMooreSearch("AFSTPABXABAF", "TPABXAB")
	fmt.Println(idx)
}
