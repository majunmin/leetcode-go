package leetcode_0560

// https://leetcode.cn/problems/subarray-sum-equals-k/?envType=study-plan-v2&envId=top-100-liked
func subarraySum(nums []int, k int) int {
	//使用前缀和的方法可以解决这个问题,因为我们需要找到和为k的连续子数组的个数. 通过计算前缀和,我们可以将问题转化为求解两个前缀和之差等于k的情况.
	//假设数组的前缀和数组为`prefixSum`,其中prefixSum[i]表示从数组起始位置到第i个位置的元素之和.
	//	那么对于任意的两个下标i和j(i < j),如果prefixSum[j] - prefixSum[i] = k，即从第i个位置到第j个位置的元素之和等于k，那么说明从第i+1个位置到第j个位置的连续子数组的和为k。
	//通过遍历数组, 计算每个位置的前缀和,并使用一个哈希表来存储每个前缀和出现的次数.在遍历的过程中，我们检查是否存在prefixSum[j] - k的前缀和,如果存在,说明从某个位置到当前位置的连续子数组的和为k,我们将对应的次数累加到结果中.
	//这样,通过遍历一次数组,我们可以统计出和为k的连续子数组的个数,并且时间复杂度为O(n),其中n为数组的长度。

	//使用前缀和.
	prefixSum := make(map[int]int)
	sum := 0
	var result int
	prefixSum[sum] = 1
	for _, num := range nums {
		sum += num
		result += prefixSum[sum-k]
		prefixSum[sum]++
	}
	return result
}

func solution1(nums []int, k int) int {
	// param check
	curSums := make([]int, 0)
	var result int
	for _, num := range nums {
		curSums = append(curSums, num)
		// result ++ if exist k - num
		for _, item := range curSums {
			if item == k-num {
				result++
			}
		}
		// add curSumMap
		for i, item := range curSums {
			curSums[i] = item + num
		}
	}
	return result
}
