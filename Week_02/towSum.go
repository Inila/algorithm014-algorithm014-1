package main

/*
	### 1. 两数之和
	使用hashset实现
*/

func towSum(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	// key 保存 数组的值, value 保存对应下标
	m := make(map[int]int)
	for idx, value := range nums {
		if valueIdx, ok := m[target-value]; ok {
			return []int{idx, valueIdx}
		}
		m[value] = idx
	}
	return []int{}
}
