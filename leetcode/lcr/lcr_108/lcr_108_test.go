package lcr_108

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},

		{
			name: "case3",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "tog", "cog"},
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				beginWord: "ymain",
				endWord:   "oecij",
				wordList:  []string{"ymann", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", "ymain"},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
