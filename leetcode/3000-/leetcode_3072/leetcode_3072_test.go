package leetcode_3072

import (
	"reflect"
	"testing"
)

func Test_resultArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{nums: []int{2, 1, 3, 3}},
			want: nil,
		},
		{
			name: "",
			args: args{nums: []int{10, 79, 12}},
			want: nil,
		},
		{
			name: "",
			args: args{nums: []int{23, 55, 90}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
