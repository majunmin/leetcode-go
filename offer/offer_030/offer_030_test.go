package offer_030

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Min())
	minStack.Push(-4)
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Min())

}
