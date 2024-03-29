package sort

import (
	"fmt"
	"testing"

	"github.com/majunmin/leetcode-go/exercise_2024/sort/merge_sort"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{[]int{1, 2, 3}},
		},
		{
			name: "case2",
			args: args{[]int{3, 2, 1}},
		},
		{
			name: "case3",
			args: args{[]int{1, 5, 3, 2, 4}},
		},
		{
			name: "case4",
			args: args{[]int{1}},
		},
		{
			name: "case5",
			args: args{[]int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge_sort.Sort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
