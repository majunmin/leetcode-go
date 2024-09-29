package leetcode

import (
	"fmt"
	"testing"
)

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{"(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		},
		{
			name: "",
			args: args{"2-1+2"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{"-042"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName1(t *testing.T) {
	finder := Constructor2()
	finder.AddNum(1)
	fmt.Println(finder.FindMedian())
	finder.AddNum(2)
	fmt.Println(finder.FindMedian())
	finder.AddNum(3)
	fmt.Println(finder.FindMedian())
}
