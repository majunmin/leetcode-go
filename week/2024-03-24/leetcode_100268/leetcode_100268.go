package leetcode_100268

import (
	"math"
	"strings"
)

type Trie struct {
	childs [26]*Trie
	idx    int
	minLen int
	isEnd  bool
}

func NewTrie() *Trie {
	return &Trie{
		childs: [26]*Trie{},
		idx:    math.MaxInt,
		minLen: math.MaxInt,
		isEnd:  false,
	}
}

func (tr *Trie) add(word string, idx int) {
	cur := tr
	for i := 0; i < len(word); i++ {
		childIdx := word[i] - 'a'
		if cur.childs[childIdx] == nil {
			cur.childs[childIdx] = NewTrie()
		}
		if len(word) < cur.childs[childIdx].minLen {
			cur.childs[childIdx].minLen = min(cur.childs[childIdx].minLen, len(word))
			cur.childs[childIdx].idx = min(cur.childs[childIdx].idx, idx)
		}
		cur = cur.childs[childIdx]
	}
	cur.isEnd = true
}

func (tr *Trie) prefix(prefix string) *Trie {
	cur := tr
	for i := 0; i < len(prefix); i++ {
		childIdx := prefix[i] - 'a'
		if cur.childs[childIdx] == nil {
			return nil
		}
		cur = cur.childs[childIdx]
	}
	return cur
}

// https://leetcode.cn/problems/longest-common-suffix-queries/
func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	// build trie
	trie := NewTrie()
	var minIdx int
	for i, word := range wordsContainer {
		if len(word) < len(wordsContainer[minIdx]) {
			minIdx = i
		}
		trie.add(reverse(word), i)
	}
	result := make([]int, 0, len(wordsQuery))
	for _, word := range wordsQuery {
		reverseWord := reverse(word)
		var cur *Trie
		for i := len(reverseWord); i > 0; i-- {
			suffix := reverseWord[:i]
			node := trie.prefix(suffix)
			if node != nil {
				cur = node
				break
			}
		}
		if cur == nil {
			result = append(result, minIdx)
		} else {
			result = append(result, cur.idx)
		}
	}
	return result
}

func reverse(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
