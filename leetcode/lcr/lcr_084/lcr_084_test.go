package lcr_084

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{[]int{1, 1, 2}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
