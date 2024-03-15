package leetcode_100251

import (
	"reflect"
	"testing"
)

func Test_findAllSubStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case",
			args: args{"abcd"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllSubStr(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
