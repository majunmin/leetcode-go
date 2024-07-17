package lintcode_0092

import "testing"

func TestBackPack(t *testing.T) {
	type args struct {
		m int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				m: 10,
				a: []int{3, 4, 8, 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackPack(tt.args.m, tt.args.a); got != tt.want {
				t.Errorf("BackPack() = %v, want %v", got, tt.want)
			}
		})
	}
}
