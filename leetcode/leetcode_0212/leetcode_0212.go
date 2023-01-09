package leetcode_0212

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func findWords(board [][]byte, words []string) []string {
	return dfsSolution2(board, words)
}

// 执行超时
func dfsSolution2(board [][]byte, words []string) []string {
	rowSize, colSize := len(board), len(board[0])
	// param check
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}

	trie := NewTrieTree()
	for _, word := range words {
		trie.AddWord(word)
	}

	result := make(map[string]struct{})
	visited := make([][]bool, rowSize)
	for i := 0; i < rowSize; i++ {
		visited[i] = make([]bool, colSize)
	}
	// dfs Solution
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			dfs2(board, trie, result, string(board[i][j]), i, j, visited, rowSize, colSize)
		}
	}
	res := make([]string, 0, len(result))
	for word := range result {
		res = append(res, word)
	}
	return res
}

func dfs2(board [][]byte, trie *TrieTree, result map[string]struct{}, path string,
	row int, col int, visited [][]bool, rowSize int, colSize int) {

	if trie.Contains(path) {
		result[path] = struct{}{}
	}

	// add result
	if !trie.StartsWith(path) {
		return
	}

	visited[row][col] = true
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow < 0 || newRow >= rowSize || newCol < 0 || newCol >= colSize || visited[newRow][newCol] {
			continue
		}
		path += string(board[newRow][newCol])
		dfs2(board, trie, result, path, newRow, newCol, visited, rowSize, colSize)
		path = path[:len(path)-1]

	}
	visited[row][col] = false
}

// 执行超时
func dfsSolution1(board [][]byte, words []string) []string {
	rowSize, colSize := len(board), len(board[0])
	// param check
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}

	wordSet := make(map[string]struct{}, len(words))
	for _, word := range words {
		wordSet[word] = struct{}{}
	}

	result := make(map[string]struct{})
	visited := make([]bool, rowSize*colSize)
	// dfs Solution
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			dfs1(board, wordSet, result, string(board[i][j]), i, j, visited, rowSize, colSize)
		}
	}
	res := make([]string, 0, len(result))
	for word := range result {
		res = append(res, word)
	}
	return res
}

func dfs1(board [][]byte, wordSet map[string]struct{}, result map[string]struct{}, path string,
	row int, col int, visited []bool, rowSize int, colSize int) {

	byteIdx := row*colSize + col
	if visited[byteIdx] {
		return
	}

	// add result
	if _, exist := wordSet[path]; exist {
		result[path] = struct{}{}
	}

	visited[byteIdx] = true
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow < 0 || newRow >= rowSize || newCol < 0 || newCol >= colSize {
			continue
		}
		path += string(board[newRow][newCol])
		dfs1(board, wordSet, result, path, newRow, newCol, visited, rowSize, colSize)
		path = path[:len(path)-1]

	}
	visited[byteIdx] = false
}

type TrieTree struct {
	buckets [26]*TrieTree
	isEnd   bool
}

func NewTrieTree() *TrieTree {
	return &TrieTree{}
}

func (t *TrieTree) AddWord(word string) {
	curTrie := t
	for _, b := range word {
		curIdx := b - 'a'
		if curTrie.buckets[curIdx] == nil {
			curTrie.buckets[curIdx] = NewTrieTree()
		}
		curTrie = curTrie.buckets[curIdx]
	}
	curTrie.isEnd = true
}

func (t *TrieTree) StartsWith(prefix string) bool {
	curTrie := t
	for _, b := range prefix {
		curTrie = curTrie.buckets[b-'a']
		if curTrie == nil {
			return false
		}
	}
	return true
}

func (t *TrieTree) Contains(word string) bool {
	curTrie := t
	for _, b := range word {
		curTrie = curTrie.buckets[b-'a']
		if curTrie == nil {
			return false
		}
	}
	return curTrie.isEnd
}

func (t *TrieTree) HasChildren() bool {
	for _, b := range t.buckets {
		if b != nil {
			return true
		}
	}
	return false
}
