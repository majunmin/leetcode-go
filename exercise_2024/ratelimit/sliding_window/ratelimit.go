package sliding_window

import "time"

type RateLimit struct {
	windowSize int   // 窗口大小,ms
	maxReqs    int   // 窗口限流阈值
	total      int   // 窗口内 总的请求计数
	shardCnt   int   // 窗口分片数(滑动窗口)
	shardSize  int   // 窗口大小(ms)
	shardId    int   // 窗口 分片id
	shards     []int // 每个分片内请求计数
	curWindow  int   // ms
}

func NewRateLimit(windowSize int, shardCnt int, maxReqs int) *RateLimit {
	return &RateLimit{
		windowSize: windowSize,
		shardCnt:   shardCnt,
		shards:     make([]int, shardCnt),
		shardSize:  windowSize / shardCnt,
		total:      maxReqs,
		curWindow:  int(time.Now().UnixMilli()),
	}
}

func (rl *RateLimit) tryAcquire() bool {
	now := int(time.Now().UnixMilli())

	// 滑动窗口
	for now > rl.curWindow {
		rl.shardId = (rl.shardId + 1) % rl.shardCnt
		rl.total -= rl.shards[rl.shardId]
		rl.shards[rl.shardId] = 0
		rl.curWindow += rl.shardSize
	}
	// 判断是否达到限流阈值
	if rl.total >= rl.maxReqs {
		return false
	}
	rl.total += 1
	rl.shards[rl.shardId]++
	return true
}
