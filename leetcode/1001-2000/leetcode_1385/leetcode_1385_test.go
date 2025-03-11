package leetcode_1385

import "testing"

func Test_findTheDistanceValue(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		d    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr1: []int{1, 2, 3, 100},
				arr2: []int{-5, -3, -2, 7, 10},
				d:    6,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDistanceValue(tt.args.arr1, tt.args.arr2, tt.args.d); got != tt.want {
				t.Errorf("findTheDistanceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
