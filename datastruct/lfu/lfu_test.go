// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-22 00:44
package lfu

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfuCache := Constructor(2)

	lfuCache.Put(1, 1)
	lfuCache.Put(2, 2)
	fmt.Println(lfuCache.Get(1))
	lfuCache.Put(3, 3)
	fmt.Println(lfuCache.Get(2))
	fmt.Println(lfuCache.Get(3))
	lfuCache.Put(4, 4)
	fmt.Println(lfuCache.Get(1))
	fmt.Println(lfuCache.Get(3))
	fmt.Println(lfuCache.Get(4))

	//lfuCache := Constructor(0)
	//lfuCache.Put(0, 0)
	//fmt.Println(lfuCache.Get(0))
}
