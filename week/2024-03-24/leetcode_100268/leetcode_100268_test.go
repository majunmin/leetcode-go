package leetcode_100268

import (
	"reflect"
	"testing"
)

func Test_stringIndices(t *testing.T) {
	type args struct {
		wordsContainer []string
		wordsQuery     []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				wordsContainer: []string{"abcd", "bcd", "xbcd"},
				wordsQuery:     []string{"xyz"},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				wordsContainer: []string{"abcdefgh", "poiuygh", "ghghgh"},
				wordsQuery:     []string{"acbfgh"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringIndices(tt.args.wordsContainer, tt.args.wordsQuery); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
