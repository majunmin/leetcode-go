package lcr_008

import "testing"

func Test_minSubArrayLen2(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen2(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen2() = %v, want %v", got, tt.want)
			}
		})
	}
}
