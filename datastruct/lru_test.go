// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-21 00:17
package datastruct

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {

	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	fmt.Println(lru.Get(1))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
}
