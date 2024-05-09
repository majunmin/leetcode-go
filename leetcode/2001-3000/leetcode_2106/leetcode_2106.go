package leetcode_2106

// https://leetcode.cn/problems/maximum-fruits-harvested-after-at-most-k-steps/
func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	// param check
	//
	var (
		preSum  = make([]int, len(fruits)+1)
		indices = make([]int, len(fruits))
		result  int
	)
	// 构建前缀和数组,  方便计算子数组和
	for i := 0; i < len(fruits); i++ {
		preSum[i+1] = preSum[i] + fruits[i][1]
		indices[i] = fruits[i][0]
	}
	// 一共走 k 步
	//1. 先向左走x 步, 向右走y = (k-2x)步， 此时边界为  [startPos - x, startPos + k -2x]
	//2. 先向右走x 步, 向左走y = (k-2x)步， 此时边界为  [startPos - k + 2x, startPos + x]
	for x := 0; 2*x <= k; x++ {
		//1  先向左走x步, 再向右边走 k - 2x步
		y := k - 2*x
		l, r := startPos-x, startPos+y
		pl := lowerBound(indices, 0, len(fruits)-1, l)
		pr := upperBound(indices, 0, len(fruits)-1, r)
		result = max(result, preSum[pr]-preSum[pl])

		l, r = startPos-y, startPos+x
		pl = lowerBound(indices, 0, len(fruits)-1, l)
		pr = upperBound(indices, 0, len(fruits)-1, r)
		result = max(result, preSum[pr]-preSum[pl])
	}
	return result
}

func upperBound(indices []int, l int, r int, target int) int {
	//for l < r {
	//	mid := l + (r-l+1)>>1
	//	if target >= indices[mid] {
	//		l = mid
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return l + 1
	res := r + 1
	for l <= r {
		mid := l + (r-l)/2
		if indices[mid] > target {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

// 找到左边  >= target 的第一个索引
func lowerBound(indices []int, l int, r int, target int) int {
	//for l < r {
	//	mid := l + (r-l)>>1
	//	if target <= indices[mid] {
	//		r = mid
	//	} else {
	//		l = mid + 1
	//	}
	//}
	//return l

	res := r + 1
	for l <= r {
		mid := l + (r-l)/2
		if indices[mid] >= target {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res

}
