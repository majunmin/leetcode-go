package leetcode_0460

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lfuCache := Constructor(2)
	lfuCache.Put(1, 1)
	lfuCache.Put(2, 2)
	fmt.Print(lfuCache.Get(1))
	lfuCache.Put(3, 3)
	fmt.Print(lfuCache.Get(2))
	fmt.Print(lfuCache.Get(3))
	lfuCache.Put(4, 4)
	fmt.Print(lfuCache.Get(1))
	fmt.Print(lfuCache.Get(3))
	fmt.Print(lfuCache.Get(4))

}
