package leetcode_0307

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumArray_SumRange(t *testing.T) {
	numArray := Constructor([]int{0, 9, 5, 7, 3})
	fmt.Println(numArray.SumRange(4, 4))
}

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
