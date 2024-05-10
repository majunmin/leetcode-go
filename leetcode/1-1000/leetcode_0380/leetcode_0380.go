package leetcode_0380

import "math/rand/v2"

type RandomizedSet struct {
	nums    []int
	indices map[int]int // 插入的val 和 在 数组中的索引 的映射
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    make([]int, 200_001),
		indices: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.indices[val]
	if exist {
		// val 已经存在, 返回 false
		return false
	}
	// 在数组中插入值
	this.nums = append(this.nums, val)
	this.indices[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exist := this.indices[val]
	if !exist {
		return false
	}
	// 移除 nums[idx]
	// 将 this.indices[len(this.indices)-1] 移除
	this.nums[idx] = this.nums[len(this.nums)-1]
	this.indices[this.nums[idx]] = idx
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.IntN(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
