package leetcode_0732

import (
	"fmt"
	"github.com/majunmin/leetcode-go/leetcode/1-1000/leetcode_0732/solution1"
	"testing"
)

func TestName(t *testing.T) {
	calendar := solution1.Constructor()
	fmt.Println(calendar.Book(1, 1_000_000_000))
	fmt.Println("end")
	//fmt.Println(calendar.Book(10, 20))
	//fmt.Println(calendar.Book(50, 60))
	//fmt.Println(calendar.Book(10, 40))
	//fmt.Println(calendar.Book(5, 15))
	//fmt.Println(calendar.Book(5, 10))
	//fmt.Println(calendar.Book(25, 55))
}
