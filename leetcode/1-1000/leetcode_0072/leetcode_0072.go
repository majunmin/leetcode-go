package leetcode_0072

// https://leetcode.cn/problems/edit-distance/
// 题解:https://leetcode.cn/problems/edit-distance/solutions/188814/dong-tai-gui-hua-java-by-liweiwei1419/
//
// 题目是  给出最短编辑距离, 而不是  给出所有遍历结果
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	op := make([][]int, 0, l1+1)
	for i := 0; i <= l1; i++ {
		op = append(op, make([]int, l2+1))
	}
	//
	for i := 1; i <= l1; i++ {
		op[i][0] = i
	}
	for i := 1; i <= l2; i++ {
		op[0][i] = i
	}
	// initstate
	//  |     |  #  | r   | o   | s   |
	//  | :-- | :-- | :-- | :-- | :-- |
	//  |  #  | 0   | 1   | 2   | 3   |
	//  | h   | 1   |     |     |     |
	//  | o   | 2   |     |     |     |
	//  | r   | 3   |     |     |     |
	//  | s   | 4   |     |     |     |
	//  | e   | 5   |     |     |     |
	//  |     |     |     |     |     |
	// 状态 转移方程
	// horse ros
	//
	// dp[i][j] = min ( dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1] + 1) // word1[i] != word2[j]
	// dp[i][j] = min ( dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1])     // word1[i] == word2[j]
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			min := op[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				min += 1
			}
			min = minInt(min, op[i-1][j]+1)
			min = minInt(min, op[i][j-1]+1)
			op[i][j] = min
		}
	}
	return op[l1][l2]

}

// dsf solution 解法超时
func dfsSolution(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	if l1 == 0 || l2 == 0 {
		return maxInt(l1, l2)
	}

	if word1[l1-1] == word2[l2-1] {
		return minDistance(word1[:l1-1], word2[:l2-1])
	}

	// dfs
	min := minDistance(word1[:l1-1], word2[:l2-1])
	min = minInt(min, minDistance(word1[:l1-1], word2))
	min = minInt(min, minDistance(word1, word2[:l2-1]))
	return min + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
