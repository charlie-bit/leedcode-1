package Leedcode_1

/**
Author:charlie
Description:
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 */
func TwoSum(nums []int, target int) []int {
	var resp = make(map[int]int)
	for key, value := range nums {
		if _,ok := resp[target-value];ok {
			return []int{resp[target-value],key}
		}else {
			resp[value] = key
		}
	}
	return nil
}