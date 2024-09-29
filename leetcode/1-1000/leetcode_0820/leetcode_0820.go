package leetcode_0820

import "slices"

// https://leetcode.cn/problems/short-encoding-of-words/_0820
func minimumLengthEncoding(words []string) int {
	// 按长度从大到小排序
	slices.SortFunc(words, func(a, b string) int {
		return len(b) - len(a)
	})
	var (
		trieTree = newTrie()
		result   int
	)
	for _, word := range words {
		if trieTree.startsWith(word) {
			continue
		}
		result += len(word) + 1
		trieTree.insert(word)
	}
	return result
}

type trie struct {
	childs [26]*trie
	isEnd  bool
}

func newTrie() *trie {
	return &trie{
		childs: [26]*trie{},
		isEnd:  false,
	}
}

// 倒序插入字符
func (t *trie) insert(word string) {
	cur := t
	for i := len(word); i >= 0; i-- {
		idx := word[i] - 'a'
		if cur.childs[idx] == nil {
			cur.childs[idx] = newTrie()
		}
		cur = cur.childs[idx]
	}
	cur.isEnd = true
}

func (t *trie) startsWith(prefix string) bool {
	cur := t
	for i := len(prefix) - 1; i >= 0; i-- {
		if cur.childs[prefix[i]-'a'] == nil {
			return false
		}
		cur = cur.childs[prefix[i]-'a']
	}
	return true
}
