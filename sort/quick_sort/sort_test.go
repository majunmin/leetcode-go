package quick_sort

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{arr: []int{0}}},
		{name: "case2", args: args{arr: []int{0, 3, 1}}},
		{name: "case3", args: args{arr: []int{0, 4, 68, 1, 3, 9, 10, 4, 5, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
