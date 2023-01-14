package leetcode_0641

import (
	"fmt"
	"testing"
)

func Test_MyCircularDeque(t *testing.T) {
	cDeque := Constructor(3)
	fmt.Println(cDeque.InsertLast(1))
	fmt.Println(cDeque.InsertLast(2))
	fmt.Println(cDeque.InsertFront(3))
	fmt.Println(cDeque.InsertFront(4))
	fmt.Println(cDeque.GetRear())
	fmt.Println(cDeque.IsFull())
	fmt.Println(cDeque.DeleteLast())
	fmt.Println(cDeque.InsertFront(4))
	fmt.Println(cDeque.GetFront())
}
