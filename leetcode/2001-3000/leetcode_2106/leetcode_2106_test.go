package leetcode_2106

import "testing"

func Test_maxTotalFruits(t *testing.T) {
	type args struct {
		fruits   [][]int
		startPos int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				fruits: [][]int{
					{0, 9},
					{4, 1},
					{5, 7},
					{6, 2},
					{7, 4},
					{10, 9},
				},
				startPos: 5,
				k:        4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalFruits(tt.args.fruits, tt.args.startPos, tt.args.k); got != tt.want {
				t.Errorf("maxTotalFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upperBound(t *testing.T) {
	type args struct {
		indices []int
		l       int
		r       int
		target  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				indices: []int{1, 3, 5, 7, 9}, l: 0, r: 4, target: 0,
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				indices: []int{1, 3, 5, 7, 9}, l: 0, r: 4, target: 5,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				indices: []int{1, 3, 5, 5, 7, 9}, l: 0, r: 4, target: 5,
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				indices: []int{1, 3, 5, 7, 9}, l: 0, r: 4, target: 11,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upperBound(tt.args.indices, tt.args.l, tt.args.r, tt.args.target); got != tt.want {
				t.Errorf("upperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowerBound(t *testing.T) {
	type args struct {
		indices []int
		l       int
		r       int
		target  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				indices: []int{1, 3, 5, 7, 9}, l: 0, r: 4, target: 0,
			},
			want: 0,
		},
		{
			name: "case2",
			args: args{
				indices: []int{1, 3, 5, 5, 7, 9}, l: 0, r: 4, target: 5,
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				indices: []int{1, 3, 5, 7, 9}, l: 0, r: 4, target: 11,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowerBound(tt.args.indices, tt.args.l, tt.args.r, tt.args.target); got != tt.want {
				t.Errorf("lowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
