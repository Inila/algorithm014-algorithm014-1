package main

import (
	"fmt"

	"github.com/go-playground/locales/ar"
	// "math"
)

/*
	11. 盛最多的水的容器：https://leetcode-cn.com/problems/container-with-most-water/。

	思路1: 枚举,即将所有的可能性列举出来进行比较, 时间复杂度为 O(n^2), 因为需要嵌套遍历数据

	思路2: 长度和高度两个变量，先选定一个最大的，即长度最长(数组下标 0 和 length -1)，然后将下标向中间移动直到相遇，每次一定高度小的下标
*/

func func1(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	// 记录当前最大面积
	var maxArea int
	/* 这里要熟练掌握嵌套循环数据遍历方法 */
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			// 面积需要按照柱子矮的来计算，所以这里计算两个的最小值
			min := nums[i]
			if nums[j] < min {
				min = nums[j]
			}

			area := (j - i) * min

			if area > maxArea {
				maxArea = area
			}
		}
	}
	fmt.Printf("maxArea: %d\n", maxArea)
}

func func2(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	/* 这里要熟练掌握数组从两边向中间遍历的方法 */
	maxArea := 0

	i, j := 0, length-1
	for i < j {
		min := nums[i]
		if nums[j] < min {
			min = nums[j]
			j--
		} else {
			i++
		}

		area := min * (j - i + 1)

		if maxArea < area {
			maxArea = area
		}
	}
	fmt.Printf("maxArea: %d\n", maxArea)
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	func1(nums)
	func2(nums)
}

/*======== 第二刷 ==========*/
// 暴力解法: 枚举所有可能性
func func152(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	maxArea := 0
	for i := 0; i < length - 1;i++ {
		for j := i+; j < length; j++ {
			// 取柱子比较矮的
			min := nums[i]
			if nums[j] < min {
				min = nums[j]
			}

			area := (j - i) * min
			if maxArea < area {
				maxmaxArea = area
			}
		}
	}
}

// 两遍夹逼解法
func func252(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	i, j := 0, length-1
	maxArea := 0
	for i < j {
		min := nums[i]
		if nums[j] < min {
			min = nums[j]
			j--
		} else {
			i++
		}

		area := (j - i + 1) * min
		if maxArea < area {
			maxArea = area
		}
	}

}

/*======== 第三刷 ==========*/
