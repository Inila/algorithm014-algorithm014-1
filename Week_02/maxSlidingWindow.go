package main

import "container/heap"

/*
	1. 滑动窗口的最大值
	给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
*/

// 暴力解法
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 0 || k <= 0 || k > len(nums) {
		return nil
	}

	maxInt := make([]int, 0, k)
	for i := 0; i < len(nums)-k; i++ {
		// maxInt = append(maxInt, findMaxByDoublePoint(nums[i:i+k]))
		maxInt = append(maxInt, findMaxByHeap(nums[i:i+k]))
	}
	return maxInt
}

func findMaxByDoublePoint(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left] > nums[right] {
			right--
		} else {
			left++
		}
	}
	return nums[left]
}

func findMaxByHeap(nums []int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	return heap.Pop(&h).(int)
}
