package leetcode_0707

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	myList := Constructor()
	myList.AddAtHead(2)
	myList.DeleteAtIndex(1)
	myList.AddAtHead(2)
	myList.AddAtHead(7)
	myList.AddAtHead(3)
	myList.AddAtHead(2)
	myList.AddAtHead(5)
	myList.AddAtTail(5)
	myList.AddAtIndex(3, 0)
	myList.DeleteAtIndex(2)
	myList.AddAtHead(6)
	myList.AddAtTail(4)
	fmt.Println(myList.Get(4))
	myList.AddAtHead(4)
	myList.AddAtIndex(5, 0)
	myList.AddAtHead(6)

}
