package offer_058

import (
	"fmt"
	"testing"
)

func TestReverseWord(t *testing.T) {
	fmt.Println(reverseWords("hello world"))
	fmt.Println(reverseWords("  hello  world  "))
	fmt.Println(reverseWords("hello"))
	fmt.Println(reverseWords("hello  "))
}
