package leetcode_0131

import (
	"reflect"
	"testing"
)

func Test_partition2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case1",
			args: args{s: "aab"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition2() = %v, want %v", got, tt.want)
			}
		})
	}
}
