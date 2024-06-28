package leetcode_100267

import "testing"

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		coins []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				coins: []int{5, 2},
				k:     7,
			},
			want: 12,
		},
		{
			name: "",
			args: args{
				coins: []int{10, 1, 3, 4, 8},
				k:     7,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest(tt.args.coins, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
