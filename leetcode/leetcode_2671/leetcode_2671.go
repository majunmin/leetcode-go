package leetcode_2671

// https://leetcode.cn/problems/frequency-tracker/?envType=daily-question&envId=2024-03-21
type FrequencyTracker struct {
	wordFreq map[int]int // 记录num出现的频次
	freqSet  map[int]int // 统计每个频次出现的次数
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		wordFreq: make(map[int]int),
		freqSet:  make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	cur := this.wordFreq[number]
	this.wordFreq[number]++
	if cur > 0 {
		this.freqSet[cur]--
	}
	this.freqSet[cur+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	freq := this.wordFreq[number]
	if freq == 0 {
		return
	}
	this.wordFreq[number]--
	if this.wordFreq[number] > 0 {
		this.freqSet[freq-1]++
	}
	this.freqSet[freq]--
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freqSet[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
