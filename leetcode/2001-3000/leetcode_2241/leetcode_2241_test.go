package leetcode_2241

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	atm := Constructor()
	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(atm.Withdraw(600))
	atm.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println(atm.Withdraw(600))
	fmt.Println(atm.Withdraw(550))
}
