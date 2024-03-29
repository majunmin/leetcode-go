package datastruct

type Trie struct {
	childs [26]*Trie
	isEnd  bool
}

func Constructor() Trie {
	return Trie{
		childs: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := range word {
		if cur.childs[word[i]-'a'] == nil {
			node := Constructor()
			cur.childs[word[i]-'a'] = &node
		}
		cur = cur.childs[word[i]-'a']
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := range word {
		if cur.childs[word[i]-'a'] == nil {
			return false
		}
		cur = cur.childs[word[i]-'a']
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := range prefix {
		if cur.childs[prefix[i]-'a'] == nil {
			return false
		}
		cur = cur.childs[prefix[i]-'a']
	}
	return true
}
