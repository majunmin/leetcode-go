package leak_bucket

import "time"

type RateLimiter struct {
	capacity      int
	water         int // 桶中剩余水量
	leakTimestamp int
	leakRate      int // 漏水频率, per second
}

func NewRateLimiter(capacity int, leakRate int) *RateLimiter {
	return &RateLimiter{
		capacity: capacity,
		leakRate: leakRate,
	}
}

// 桶中的水满了， 拒绝请求
// 否则放行
func (rl *RateLimiter) tryAcquire() bool {
	// 如果 桶中没有水
	if rl.water == 0 {
		rl.leakTimestamp = int(time.Now().UnixMilli())
		rl.water++
		return rl.water < rl.capacity
	}
	now := int(time.Now().UnixMilli())

	// 先漏水, 计算剩余水量
	leakedWater := (now - rl.leakTimestamp) / 1000 * rl.leakRate
	if leakedWater != 0 {
		leftWater := rl.water - leakedWater
		rl.water = max(0, leftWater)
		rl.leakTimestamp = now
	}

	if rl.water < rl.capacity {
		rl.water++
		return true
	}
	// 桶中的水满了
	return false
}
