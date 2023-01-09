package offer_037

import (
	"fmt"
	"testing"
)

var (
	codec = &Codec{}
)

func Test_Serialize(t *testing.T) {
	root := codec.Deserialize("[1,2,3,nil,nil,4,5]")
	fmt.Println(codec.Serialize(root))
}
