package leetcode

import "testing"

func Test_convertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{100},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
