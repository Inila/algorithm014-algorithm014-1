package main

/*
1. 两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

*/

// 暴力解法
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if target-nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 使用hashSet解法，以空间换区时间
func twoSum2(nums []int, taget int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	// key 记录遍历的值，value记录下标
	m := make(map[int]int)
	for idx, value := range nums {
		if value, ok := m[target-value]; ok {
			return []int{value, idx}
		}
		m[value] = idx
	}
	return []int{}
}

func main() {
	nums := [2, 7, 11, 15]
	target := 9
	twoSum1(nums, target)
	twoSum2(nums, target)
}
