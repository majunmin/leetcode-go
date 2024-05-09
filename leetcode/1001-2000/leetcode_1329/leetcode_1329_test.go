package leetcode_1329

import (
	"reflect"
	"testing"
)

func Test_diagonalSort(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{[][]int{
				{3, 3, 1, 1},
				{2, 2, 1, 2},
				{1, 1, 1, 2},
			}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalSort(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagonalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
