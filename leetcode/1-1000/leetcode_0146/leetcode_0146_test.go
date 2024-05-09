package leetcode_0146

import (
	"fmt"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(2))
	lru.Put(1, 1)
	lru.Put(4, 1)
	fmt.Println(lru.Get(2))
}
