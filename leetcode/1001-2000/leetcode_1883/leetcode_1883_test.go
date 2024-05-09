package leetcode_1883

import "testing"

func Test_minSkips(t *testing.T) {
	type args struct {
		dist        []int
		speed       int
		hoursBefore int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				dist:        []int{7, 3, 5, 5},
				speed:       2,
				hoursBefore: 10,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				dist:        []int{7, 3, 5, 5},
				speed:       1,
				hoursBefore: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSkips2(tt.args.dist, tt.args.speed, tt.args.hoursBefore); got != tt.want {
				t.Errorf("minSkips() = %v, want %v", got, tt.want)
			}
		})
	}
}
