package lcr_121

import "testing"

func Test_findTargetIn2DPlants(t *testing.T) {
	type args struct {
		plants [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				plants: [][]int{{-1, 3}},
				target: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetIn2DPlants(tt.args.plants, tt.args.target); got != tt.want {
				t.Errorf("findTargetIn2DPlants() = %v, want %v", got, tt.want)
			}
		})
	}
}
