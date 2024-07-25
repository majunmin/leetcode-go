package leetcode_0211

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{trie: NewTrie()}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}

type Trie struct {
	childs [26]*Trie
	isEnd  bool
}

func NewTrie() *Trie {
	return &Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.childs[idx] == nil {
			cur.childs[idx] = &Trie{}
		}
		cur = cur.childs[idx]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return true
	}
	return this.dfsSearch(word, 0)
}

func (this *Trie) dfsSearch(word string, i int) bool {
	if i == len(word) {
		return this.isEnd
	}
	if word[i] == '.' {
		for _, next := range this.childs {
			if next != nil && next.dfsSearch(word, i+1) {
				return true
			}
		}
	}
	idx := word[i] - 'a'
	return this.childs[idx] != nil && this.childs[idx].dfsSearch(word, i+1)

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
