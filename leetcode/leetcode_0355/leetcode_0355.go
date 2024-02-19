package leetcode_0355

import (
	"container/heap"
)

type hp struct {
	msgs []*Message
}

func newHp() *hp {
	return &hp{msgs: make([]*Message, 0)}
}

func (h *hp) Len() int {
	return len(h.msgs)
}

func (h *hp) Less(i, j int) bool {
	return h.msgs[i].ts > h.msgs[j].ts
}

func (h *hp) Swap(i, j int) {
	h.msgs[i], h.msgs[j] = h.msgs[j], h.msgs[i]
}

func (h *hp) Push(x any) {
	h.msgs = append(h.msgs, x.(*Message))
}

func (h *hp) Pop() any {
	item := h.msgs[len(h.msgs)-1]
	h.msgs = h.msgs[:len(h.msgs)-1]
	return item
}

type Message struct {
	id   int
	ts   int
	next *Message
}

type Twitter struct {
	// 计时器
	timestamp int
	// 维护关注关系
	followMap map[int]map[int]bool
	// userId -> 推文列表
	feeds map[int]*Message
}

func Constructor() Twitter {
	return Twitter{
		followMap: make(map[int]map[int]bool),
		feeds:     make(map[int]*Message),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.timestamp++
	// 插入 feed
	msg := &Message{
		id: tweetId,
		ts: this.timestamp,
	}
	msg.next = this.feeds[userId]
	this.feeds[userId] = msg
}

// 返回最近的10条推文
func (this *Twitter) GetNewsFeed(userId int) []int {
	// 合并有序链表
	// get feeds userIDs
	userIDs := make([]int, 0)
	userIDs = append(userIDs, userId)
	for uid := range this.followMap[userId] {
		userIDs = append(userIDs, uid)
	}
	// get top10 feeds
	p := newHp()
	heap.Init(p)
	for _, uid := range userIDs {
		if this.feeds[uid] != nil {
			heap.Push(p, this.feeds[uid])
		}
	}
	// get top10
	result := make([]int, 0, 10)
	cnt := 10
	for len(p.msgs) > 0 && cnt > 0 {
		item := heap.Pop(p).(*Message)
		result = append(result, item.id)
		cnt--
		if item.next != nil {
			heap.Push(p, item.next)
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exist := this.followMap[followerId]; !exist {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, exist := this.followMap[followerId]; !exist {
		return
	}
	delete(this.followMap[followerId], followeeId)
}
