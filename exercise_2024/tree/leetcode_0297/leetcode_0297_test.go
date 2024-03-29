package leetcode_0297

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	codec := Constructor()
	root := codec.deserialize("1,2,null,null,3,4,null,null,5,null,null")
	fmt.Println(codec.serialize(root))

	//fmt.Println(strings.Split("", ","))
	//fmt.Println(strings.Split(",", ","))
	//fmt.Println(strings.Split("aa", ","))
	//fmt.Println(strings.Split("aa,", ","))
}
