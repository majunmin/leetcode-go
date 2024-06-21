package leetcode_2806

const (
	amount = 100
)

// https://leetcode.cn/problems/account-balance-after-rounded-purchase/?envType=daily-question&envId=2024-06-12
func accountBalanceAfterPurchase(purchaseAmount int) int {
	cost := purchaseAmount / 10 * 10
	if purchaseAmount%10 >= 5 {
		cost += 10
	}
	return amount - cost
}

func accountBalanceAfterPurchase2(purchaseAmount int) int {
	return amount - (purchaseAmount+5)/10*10
}
