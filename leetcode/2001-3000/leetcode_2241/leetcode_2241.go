package leetcode_2241

// https://leetcode.cn/problems/design-an-atm-machine/

var (
	prices = []int{20, 50, 100, 200, 500}
)

type ATM struct {
	price2Cnt map[int]int
}

func Constructor() ATM {
	return ATM{price2Cnt: make(map[int]int)}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, cnt := range banknotesCount {
		this.price2Cnt[prices[i]] += cnt
	}
}

func (this *ATM) Withdraw(amount int) []int {
	if amount%10 > 0 {
		return []int{-1}
	}
	result := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		if amount >= prices[i] {
			if this.price2Cnt[prices[i]] <= 0 {
				return []int{-1}
			}
			result[i] += amount / prices[i]
			amount -= result[i] * prices[i]
		}
	}

	if amount > 0 {
		return []int{-1}
	}
	// process
	for i, cnt := range result {
		this.price2Cnt[prices[i]] -= cnt
	}
	return result
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
