package leetcode_1702

import "testing"

func Test_maximumBinaryString(t *testing.T) {
	type args struct {
		binary string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{"000110"},
			want: "111011",
		},
		{
			name: "",
			args: args{"01"},
			want: "01",
		},
		{
			name: "",
			args: args{"00"},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBinaryString(tt.args.binary); got != tt.want {
				t.Errorf("maximumBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
