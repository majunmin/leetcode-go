package leetcode_0212

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'o', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{"oath", "eat"},
			},
			want: []string{"oath", "eat"},
		},
		{
			name: "test",
			args: args{
				board: [][]byte{
					{'a', 'a'},
				},
				words: []string{"a"},
			},
			want: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	tree := NewTrieTree()
	tree.AddWord("oath")
	tree.AddWord("eat")
	fmt.Println(tree.Contains("oath"))
}
