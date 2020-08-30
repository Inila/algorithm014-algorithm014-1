package main

/*
	169. 多数元素
	给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。

	1. map记录每个元素的个数，遍历数据，当 > n/2 时，直接跳出循环
	2. 先排序，取第 n/2 个元素
*/

func majorityElement(nums []int) int {
	m := make(map[int]int)

	max := len(nums) / 2

	for _, num := range nums {
		cnt, ok := m[num]
		if !ok {
			m[num] = 1
		}
		cnt++
		m[num] = cnt
		if cnt > max {
			return num
		}
	}
	return 0
}
