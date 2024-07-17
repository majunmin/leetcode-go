package lcr_177

import (
	"reflect"
	"testing"
)

func Test_sockCollocation(t *testing.T) {
	type args struct {
		sockets []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{[]int{4, 5, 2, 4, 6, 6}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sockCollocation(tt.args.sockets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sockCollocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
