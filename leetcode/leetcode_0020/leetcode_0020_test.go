// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-19 22:49
package leetcode_0020

import (
	"fmt"
	"testing"
)

func Test_IsValid(t *testing.T) {
	fmt.Println(isValid("{{{}}}"))
	fmt.Println(isValid("(())()"))
	fmt.Println(isValid("((])()"))
	fmt.Println(isValid("}"))
	fmt.Println(isValid(""))
}
