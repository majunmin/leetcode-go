package leetcode_0676

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	magicDictionary := Constructor()
	magicDictionary.BuildDict([]string{"leetcode", "hello"})
	fmt.Println(magicDictionary.Search("hello"))
	fmt.Println(magicDictionary.Search("hellx"))
	fmt.Println(magicDictionary.Search("heelo"))
	fmt.Println(magicDictionary.Search("hell"))
}
