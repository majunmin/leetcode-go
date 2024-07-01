package lcr_0135

import "testing"

func Test_countNumbers(t *testing.T) {
	type args struct {
		cnt int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case-2",
			args: args{2},
		},
		{
			name: "case-3",
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countNumbers(tt.args.cnt)
		})
	}
}
