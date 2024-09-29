package leetcode_0676

type MagicDictionary struct {
	root *trie
}

func Constructor() MagicDictionary {
	return MagicDictionary{root: new(trie)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.root.insert(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.root.query(searchWord, 0, 1)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

type trie struct {
	childs [26]*trie
	isEnd  bool
}

func (t *trie) insert(word string) {
	cur := t
	for _, c := range word {
		if cur.childs[c-'a'] == nil {
			cur.childs[c-'a'] = new(trie)
		}
		cur = cur.childs[c-'a']
	}
	cur.isEnd = true
}

func (t *trie) search(word string) bool {
	cur := t
	for _, c := range word {
		if cur.childs[c-'a'] == nil {
			return false
		}
		cur = cur.childs[c-'a']
	}
	return cur.isEnd
}

func (t *trie) query(word string, idx, limit int) bool {
	if limit < 0 {
		return false
	}
	if idx == len(word) {
		return t.isEnd && limit == 0
	}
	// 枚举 26个字母
	for i := 'a'; i <= 'z'; i++ {
		c := byte(i)
		if t.childs[c-'a'] == nil {
			continue
		}
		nextLimit := limit
		if c != word[idx] {
			nextLimit -= 1
		}
		if t.childs[c-'a'].query(word, idx+1, nextLimit) {
			return true
		}
	}
	return false
}
