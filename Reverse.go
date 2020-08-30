package Leedcode_1

import (
	"math"
	"strconv"
)

func Reverse(x int) int {
	var resp string
	for x != 0{
		if x < 0 {
			resp += "-"
			x = -x
		}
		resp += strconv.Itoa(x % 10)
		x = x/10
	}
	result,_ := strconv.Atoi(resp)
	if result < int(math.Pow(-2,31)) || result > int(math.Pow(2,31) -1 ) {
		result = 0
	}
	return result
}
