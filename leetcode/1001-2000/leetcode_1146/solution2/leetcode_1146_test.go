package solution2

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	sArr := Constructor(2)
	fmt.Println(sArr.Snap())
	fmt.Println(sArr.Get(1, 0))
	fmt.Println(sArr.Get(0, 0))
	sArr.Set(1, 8)
	fmt.Println(sArr.Get(0, 0))
	sArr.Set(0, 7)
	fmt.Println(sArr.Get(0, 0))
	sArr.Set(0, 20)

}
