package _024_02_04

import (
	"reflect"
	"testing"
)

func Test_resultGrid(t *testing.T) {
	type args struct {
		image     [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				image: [][]int{
					{5, 6, 7, 10},
					{8, 9, 10, 10},
					{11, 12, 13, 10},
				},
				threshold: 3,
			},
			want: nil,
		},
		{
			name: "case1",
			args: args{
				image: [][]int{
					{9, 4, 1},
					{5, 9, 7},
					{9, 4, 2},
				},
				threshold: 5,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultGrid(tt.args.image, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
