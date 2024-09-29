package leetcode_1087

import (
	"reflect"
	"testing"
)

func Test_expand(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{"{a,b,c}d{d,f}"},
			want: nil,
		},
		{
			name: "case2",
			args: args{"aa{a,b,c}d{d,f}"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expand() = %v, want %v", got, tt.want)
			}
		})
	}
}
