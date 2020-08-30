package Leedcode_1

import (
	"strconv"
)

func IsPalindrome(x int) bool {

	param := strconv.Itoa(x)
	for i := 0; i < len(param); i++ {
		if param[i] != param[len(param)-1-i] {
			return false
		}
	}
	return true
}