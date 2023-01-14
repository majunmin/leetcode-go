package unifind

import (
	"fmt"
	"testing"
)

func Test_UF(t *testing.T) {
	uf := NewUniFind(7)
	fmt.Println(uf.Count())
	uf.Union(0, 1)
	fmt.Println(uf.Count())
	uf.Union(2, 3)
	fmt.Println(uf.Count())
	uf.Union(0, 2)
	fmt.Println(uf.Count())
	fmt.Println(uf.Count())
}
