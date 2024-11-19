package demo

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	st := newSegmentTree(10)
	st.insert(1)
	st.insert(2)
	st.insert(3)
	st.insert(3)
	st.insert(3)
	fmt.Println(st.count(1, 5))
	fmt.Println(st.count(1, 1))
	fmt.Println(st.count(4, 5))
	fmt.Println(st.getKth(1, 5, 1))
	fmt.Println(st.getKth(1, 5, 2))
	fmt.Println(st.getKth(1, 5, 3))
	fmt.Println(st.getKth(1, 5, 4))
}
