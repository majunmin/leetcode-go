package leetcode_0299

import "fmt"

// https://leetcode.cn/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	// param check len(secret) = len(guess)
	cntSecret, cntGuess := make([]int, 10), make([]int, 10)
	var (
		bullCnt int
		cowCnt  int
	)

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bullCnt++
		}
		cntSecret[secret[i]-'0']++
		cntGuess[guess[i]-'0']++
	}
	for i := 0; i < len(secret); i++ {
		// 异位数字
		cowCnt += min(cntSecret[i], cntGuess[i])
	}
	return fmt.Sprintf("%dA%dB", bullCnt, cowCnt)
}
