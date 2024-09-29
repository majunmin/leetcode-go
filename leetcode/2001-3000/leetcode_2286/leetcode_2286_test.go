package leetcode_2286

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	book := Constructor(2, 5)
	fmt.Println(book.Gather(4, 0))
	fmt.Println(book.Gather(2, 0))
	fmt.Println(book.Scatter(5, 1))
	fmt.Println(book.Scatter(5, 1))

	book = Constructor(2, 1)
	fmt.Println(book.Scatter(2, 1))
	fmt.Println(book.Gather(1, 1))
}
