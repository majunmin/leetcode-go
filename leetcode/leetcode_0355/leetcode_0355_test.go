package leetcode_0355

import (
	"fmt"
	"testing"
)

func TestTwiter(t *testing.T) {
	obj := Constructor()
	obj.PostTweet(1, 5)
	res1 := obj.GetNewsFeed(1)
	fmt.Println(res1)
	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	res2 := obj.GetNewsFeed(1)
	fmt.Println(res2)
	obj.Unfollow(1, 2)
	res3 := obj.GetNewsFeed(1)
	fmt.Println(res3)
}
