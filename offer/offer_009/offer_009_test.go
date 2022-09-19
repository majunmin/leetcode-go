package offer_009

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func Test_Cqueue(t *testing.T) {
	cq := Constructor()
	cq.AppendTail(1)
	cq.AppendTail(2)
	cq.AppendTail(3)
	cq.AppendTail(4)

	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	cq.AppendTail(5)
	cq.AppendTail(6)
	cq.AppendTail(7)
	cq.AppendTail(8)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
}
