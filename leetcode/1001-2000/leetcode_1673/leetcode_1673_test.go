package leetcode_1673

import (
	"reflect"
	"testing"
)

func Test_mostCompetitive(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2},
				k:    3,
			},
			want: nil,
		},
		{
			name: "case2",
			args: args{
				nums: []int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71},
				k:    24,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCompetitive(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostCompetitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
