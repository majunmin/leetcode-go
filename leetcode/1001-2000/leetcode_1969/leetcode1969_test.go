package leetcode_1969

import "testing"

func Test_minNonZeroProduct(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case2",
			args: args{4},
			want: 581202553,
		},
		{
			name: "case3",
			args: args{3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNonZeroProduct(tt.args.p); got != tt.want {
				t.Errorf("minNonZeroProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
