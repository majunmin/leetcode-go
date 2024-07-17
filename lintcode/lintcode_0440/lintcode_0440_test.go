package lintcode_0440

import "testing"

func TestBackPackIII(t *testing.T) {
	type args struct {
		a []int
		v []int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a: []int{2, 3, 5, 7},
				v: []int{1, 5, 2, 4},
				m: 15,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackPackIII(tt.args.a, tt.args.v, tt.args.m); got != tt.want {
				t.Errorf("BackPackIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
