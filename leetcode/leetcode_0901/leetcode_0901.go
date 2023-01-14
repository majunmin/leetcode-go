package leetcode_0901

// https://leetcode.cn/problems/online-stock-span/

// // solution1
// // StockSpanner  维护一个单调递减栈
// type StockSpanner struct {
// 	// 存储 价格索引
// 	prices []int
// 	stack  []int
// }
//
// func Constructor() StockSpanner {
// 	return StockSpanner{
// 		prices: make([]int, 0),
// 		stack:  []int{-1}, // 辅助计算
// 	}
// }
//
// func (ss *StockSpanner) Next(price int) int {
// 	currentIndex := len(ss.prices)
// 	ss.prices = append(ss.prices, price)
//
// 	for len(ss.stack) > 1 && price >= ss.prices[ss.stack[len(ss.stack)-1]] {
// 		// pop
// 		ss.stack = ss.stack[:len(ss.stack)-1]
//
// 	}
// 	res := currentIndex - ss.stack[len(ss.stack)-1]
// 	ss.stack = append(ss.stack, currentIndex)
// 	return res
// }

// solution2 这种方法 画个图也就出来了， 官方题解的方法,很简单
type StockSpanner struct {
	weightStack []int
	priceStack  []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		weightStack: make([]int, 0),
		priceStack:  make([]int, 0),
	}
}

func (ss *StockSpanner) Next(price int) int {
	res := 1
	for len(ss.priceStack) > 0 && price >= ss.priceStack[len(ss.priceStack)-1] {
		ss.priceStack = ss.priceStack[:len(ss.priceStack)-1]
		res += ss.weightStack[len(ss.weightStack)-1]
		ss.weightStack = ss.weightStack[:len(ss.weightStack)-1]
	}
	ss.priceStack = append(ss.priceStack, price)
	ss.weightStack = append(ss.weightStack, res)
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
