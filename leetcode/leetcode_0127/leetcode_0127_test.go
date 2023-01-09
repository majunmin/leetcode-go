package leetcode_0127

import (
	"fmt"
	"testing"
)

func Test_LadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
