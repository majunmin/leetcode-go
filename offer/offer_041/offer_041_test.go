package offer_041

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	finder := Constructor()
	finder.AddNum(2)
	finder.AddNum(1)
	finder.AddNum(3)
	fmt.Println(finder.arr)
	fmt.Println(finder.length)
	fmt.Println(finder.FindMedian())
}
