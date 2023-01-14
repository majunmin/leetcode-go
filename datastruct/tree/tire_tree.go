package tree

type Trie struct {
	// a-z
	bucket [26]*Trie
	isEnd  bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	//if len(word) == 0 {
	//	return
	//}
	curTrie := this
	idx := int32(0)
	for _, ch := range word {
		idx = ch - 'a'
		if curTrie.bucket[idx] == nil {
			child := Constructor()
			curTrie.bucket[idx] = &child
		}
		curTrie = curTrie.bucket[idx]
	}
	curTrie.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curTrie := this
	for _, ch := range word {
		curTrie = curTrie.bucket[ch-'a']
		//
		if curTrie == nil {
			return false
		}
	}
	return curTrie.isEnd

}

func (this *Trie) StartsWith(prefix string) bool {
	curTrie := this
	for _, ch := range prefix {
		curTrie = curTrie.bucket[ch-'a']
		if curTrie == nil {
			return false
		}
	}
	return true
}
