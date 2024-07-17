package leetcode_0130

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
		},
		{
			name: "",
			args: args{[][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			}},
		},
		{
			name: "",
			args: args{[][]byte{
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
		})
	}
}
