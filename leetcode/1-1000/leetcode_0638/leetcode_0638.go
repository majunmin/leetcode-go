package leetcode_0638

// https://leetcode.cn/problems/shopping-offers/
func shoppingOffers(price []int, special [][]int, needs []int) int {
	// len(prices) = n
	var (
		n             = len(price)
		filterSpecial = make([][]int, 0, len(special))
	)
	// 提前过滤到一些不符合条件的大礼包
	for _, s := range special {
		var totalPrice, totalCnt int
		for i := range s[:n] {
			totalCnt += s[i]
			totalPrice += s[i] * price[i]
		}
		if totalCnt > 0 && s[n] < totalPrice {
			filterSpecial = append(filterSpecial, s)
		}
	}
	// dfs
	var (
		memo = make(map[string]int) // curNeeds => minPrice
		dfs  func(curNeeds []byte) int
	)

	dfs = func(curNeeds []byte) int {
		// terminate
		if _, exist := memo[string(curNeeds)]; exist {
			return memo[string(curNeeds)]
		}
		//计算不使用大礼包的 minPrice
		var minPrice int
		for i, cnt := range curNeeds {
			minPrice += int(cnt) * price[i]
		}
		// for choice in choiceList
		nextNeeds := make([]byte, n)
		for _, s := range filterSpecial {
			satify := true
			for i, need := range curNeeds {
				if int(need) < s[n] {
					satify = false
					break
				}
				nextNeeds[i] = need - byte(s[i])
			}
			if satify {

				minPrice = min(minPrice, dfs(nextNeeds))
			}
		}
		memo[string(curNeeds)] = minPrice
		return minPrice
	}
	curNeeds := make([]byte, n)
	for i := range needs {
		curNeeds[i] = byte(needs[i])
	}
	return dfs(curNeeds)
}
