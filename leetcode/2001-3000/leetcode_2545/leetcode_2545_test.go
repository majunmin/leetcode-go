package leetcode_2545

import (
	"reflect"
	"testing"
)

func Test_sortTheStudents(t *testing.T) {
	type args struct {
		score [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				score: [][]int{
					{4, 8, 3, 15},
					{7, 5, 11, 2},
					{10, 6, 9, 1},
				},
				k: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTheStudents(tt.args.score, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTheStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
