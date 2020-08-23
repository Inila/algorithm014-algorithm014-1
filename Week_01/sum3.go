package main

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

// 求三数之和等于0,那么可以转换为 a + b = -c 即两数之和的问题，基于前面的两个之和解法，可以使用hashSet，记录所有的 -c 情况
// 先排序然后通过双指针进行左右夹逼办法进行处理
func threeSum1(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 target
	for first := 0; first < length; first++ {
		// 和上一次不能重复
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 最右端的指针
		third := length - 1
		target := -1 * nums[first]

		// 从左开始遍历
		for second := first + 1; second < length; second++ {
			// 需要和上次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 两数之和大于target，需要减小，因为有序，所以向前移动最后的指针
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 重合则不可能满足条件则退出
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
