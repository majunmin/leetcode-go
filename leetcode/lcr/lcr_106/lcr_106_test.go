package lcr_106

import "testing"

func Test_isBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{[][]int{
				{4, 1},
				{0, 2},
				{1, 3},
				{2, 4},
				{3, 0},
			},
			},
			want: false,
		},
		{
			name: "",
			args: args{[][]int{
				{1, 3},
				{0, 2},
				{1, 3},
				{0, 2},
			},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartiteDFS(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
