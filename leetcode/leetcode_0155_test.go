package leetcode

import (
	"fmt"
	"testing"
)

func TestNam(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-1)
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
