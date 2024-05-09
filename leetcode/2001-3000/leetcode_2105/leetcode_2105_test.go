package leetcode_2105

import "testing"

func Test_minimumRefill2(t *testing.T) {
	type args struct {
		plants    []int
		capacityA int
		capacityB int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				plants:    []int{2, 1, 1},
				capacityA: 2,
				capacityB: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRefill2(tt.args.plants, tt.args.capacityA, tt.args.capacityB); got != tt.want {
				t.Errorf("minimumRefill2() = %v, want %v", got, tt.want)
			}
		})
	}
}
