// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-10 23:46
package leetcode_0306

import (
	"fmt"
	"testing"
)

func TestStringAdd(t *testing.T) {
	fmt.Println(stringAdd("111", "3991"))
}

func TestValidStringAdditive(t *testing.T) {
	fmt.Println(isAdditiveNumber("23581321"))
	fmt.Println(isAdditiveNumber("1199200399"))
	fmt.Println(isAdditiveNumber("1199200400"))
}
