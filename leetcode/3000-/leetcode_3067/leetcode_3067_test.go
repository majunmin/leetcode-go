package leetcode_3067

import (
	"reflect"
	"testing"
)

func Test_countPairsOfConnectableServers(t *testing.T) {
	type args struct {
		edges       [][]int
		signalSpeed int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				edges: [][]int{
					{0, 1, 1},
					{1, 2, 5},
					{2, 3, 13},
					{3, 4, 9},
					{4, 5, 2},
				},
				signalSpeed: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairsOfConnectableServers(tt.args.edges, tt.args.signalSpeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPairsOfConnectableServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
