package leetcode_0295

import (
	"fmt"
	"testing"
)

func Test_newMinHeap(t *testing.T) {
	mediaFinder := Constructor()
	mediaFinder.AddNum(1)
	fmt.Print(mediaFinder.FindMedian())
	mediaFinder.AddNum(2)
	fmt.Print(mediaFinder.FindMedian())
	mediaFinder.AddNum(3)
	fmt.Print(mediaFinder.FindMedian())
}
