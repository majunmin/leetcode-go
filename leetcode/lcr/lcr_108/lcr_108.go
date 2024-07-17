package lcr_108

import (
	"math"
)

// https://leetcode.cn/problems/om3reC/
// 单词接龙I
func ladderLength(beginWord string, endWord string, wordList []string) int {
	return bfsSolution2(beginWord, endWord, wordList)
}

// 双向BFS
func bfsSolution2(beginWord string, endWord string, wordList []string) int {
	var (
		leftQueue  = make([]string, 0)
		rightQueue = make([]string, 0)
		visited    = make(map[string]bool)
		wordSet    = make(map[string]bool, 0)
	)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	leftQueue = append(leftQueue, beginWord)
	rightQueue = append(rightQueue, endWord)
	visited[beginWord] = true
	visited[endWord] = true
	result := 1
	for len(leftQueue) > 0 && len(rightQueue) > 0 {
		if len(rightQueue) < len(leftQueue) {
			leftQueue, rightQueue = rightQueue, leftQueue
		}
		size := len(leftQueue)
		for i := 0; i < size; i++ {
			item := leftQueue[i]
			for j := 0; j < len(item); j++ {
				itemBytes := []byte(item)
				for c := 'a'; c <= 'z'; c++ {
					itemBytes[j] = byte(c)
					next := string(itemBytes)
					if !wordSet[next] {
						continue
					}
					if inList(rightQueue, next) {
						return result + 1
					}
					if visited[next] {
						continue
					}
					leftQueue = append(leftQueue, next)
					visited[next] = true
				}

			}
		}
		result++
		leftQueue = leftQueue[size:]
	}
	return 0
}

func inList(queue []string, next string) bool {
	for _, item := range queue {
		if item == next {
			return true
		}
	}
	return false
}

func bfsSolution(beginWord string, endWord string, wordList []string) int {
	var (
		queue   = make([]string, 0)
		visited = make(map[string]bool)
		wordSet = make(map[string]bool, 0)
		result  = 1
	)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	queue = append(queue, beginWord)
	visited[beginWord] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			for j := 0; j < len(item); j++ {
				itemBytes := []byte(item)
				for c := 'a'; c <= 'z'; c++ {
					itemBytes[j] = byte(c)
					next := string(itemBytes)
					if !wordSet[next] || visited[next] {
						continue
					}
					visited[next] = true
					if next == endWord {
						return result + 1
					}
					queue = append(queue, next)
				}
			}
		}
		result++
		queue = queue[size:]
	}
	return 0
}

func dfsSolution(beginWord string, endWord string, wordList []string) int {
	var (
		wordSet = make(map[string]bool)
		visited = make(map[string]bool)
	)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}
	wordSet[endWord] = true
	visited[beginWord] = true
	res := backtrace(beginWord, endWord, wordSet, visited)
	if res == math.MaxInt-1 {
		return 0
	}
	return res
}

func backtrace(cur string, endWord string, wordSet map[string]bool, visited map[string]bool) int {
	if cur == endWord {
		return 1
	}

	var (
		res = math.MaxInt - 1
	)
	for i := 0; i < len(cur); i++ {
		curBytes := []byte(cur)
		for j := 'a'; j <= 'z'; j++ {
			curBytes[i] = byte(j)
			next := string(curBytes)
			if !wordSet[next] || visited[next] {
				continue
			}
			visited[next] = true
			res = min(res, backtrace(next, endWord, wordSet, visited)+1)
			visited[next] = false
		}
	}
	return res
}
