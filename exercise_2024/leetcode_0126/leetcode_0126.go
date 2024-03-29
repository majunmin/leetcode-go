package leetcode_0126

// https://leetcode.cn/problems/word-ladder-ii/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// param check
	if len(wordList) == 0 {
		return nil
	}
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return nil
	}
	delete(wordSet, beginWord)
	//
	parent := make(map[string][]string)
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	steps := make(map[string]int)
	steps[beginWord] = 0
	var step int
	var flag bool
	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			for j := 0; j < len(cur); j++ {
				for k := 'a'; k <= 'z'; k++ {
					next := cur[:j] + string(k) + cur[j+1:]
					if _, exist := steps[next]; exist && steps[next] == step {
						parent[next] = append(parent[next], cur)
						continue
					}
					if wordSet[next] {
						if next == endWord {
							flag = true
						}
						parent[next] = append(parent[next], cur)
						delete(wordSet, next)
						steps[next] = step
						queue = append(queue, next)
					}
				}
			}
		}
		if flag {
			break
		}
		queue = queue[size:]
	}

	if !flag {
		return nil
	}

	// 2. 回溯算法构建路径
	var result [][]string
	backtrace(&result, parent, endWord, beginWord, []string{endWord})
	return result
}

func backtrace(result *[][]string, parent map[string][]string, curWord string, endWord string, path []string) {
	// terminate
	if curWord == endWord {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	// for choice in choiceList
	for _, next := range parent[curWord] {
		path = append([]string{next}, path...)
		backtrace(result, parent, next, endWord, path)
		path = path[1:]
	}
}
