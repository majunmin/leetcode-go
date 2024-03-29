package fixed_window

import "time"

type RateLimit struct {
	windowSize    int //  ms
	limit         int // 窗口内请求数限制
	currentWindow int // 当前窗口终止位置
	count         int //当前窗口计数
}

func NewRateLimit(windowSize int, limit int) *RateLimit {
	return &RateLimit{
		windowSize:    windowSize,
		limit:         limit,
		currentWindow: int(time.Now().UnixMilli()),
	}
}

func (rl *RateLimit) tryAcquire() bool {
	now := int(time.Now().UnixMilli())
	// 按需重置窗口
	if now > rl.currentWindow {
		rl.count = 0
	}
	for now > rl.currentWindow {
		rl.currentWindow += rl.windowSize
	}
	// 达到限流阈值
	if rl.count >= rl.limit {
		return false
	}
	rl.count++
	return true
}
