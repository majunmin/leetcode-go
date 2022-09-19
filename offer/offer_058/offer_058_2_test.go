package offer_058

import (
	"fmt"
	"testing"
)

func Test_ReverseWord(t *testing.T) {

	fmt.Println(reverseLeftWords("asdgag", 3))
	fmt.Println(reverseLeftWords("leftwordkk", 8))
	fmt.Println(reverseLeftWords("leftwordkk", 0))
	fmt.Println(reverseLeftWords("leftwordkk", 9))
}
