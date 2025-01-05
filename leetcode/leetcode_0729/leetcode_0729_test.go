package leetcode_0729

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	//[[],[47,50],[33,41],[39,45],[33,42],[25,32],[26,35],[19,25],[3,8],[8,13],[18,27]]

	myCalendar := Constructor()
	fmt.Println(myCalendar.Book(20, 29)) // true
	//fmt.Println(myCalendar.Book(13, 22)) // false
	fmt.Println(myCalendar.Book(44, 50)) // true
	fmt.Println(myCalendar.Book(1, 7))   // true
	//fmt.Println(myCalendar.Book(2, 10))  // false
	fmt.Println(myCalendar.Book(14, 20)) // true
	fmt.Println(myCalendar.Book(19, 25)) // true
}
