package main

import (
	"fmt"
)

/*
	Move Zeroes: 零移动
	给定一个数组 nums, 编写一个函数，将所有 0 元素移动到数组末尾，同时保持非零元素相对顺序
*/

/*
	5-10分钟思路：
	1. 循环遍历，非0则和上一次0元素下标交换
	2. 首选循环一遍，统计0的个数和0对应的下标信息，然后将对应下标+1的数据向前移动，最后有几个0将最后几位设置为0
	3. 申请一个新的数组，遍历数组，非0的写入到新数组，0的则放到最后。 (不符合题目要求)
*/
func func1(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
	}
	fmt.Printf("func1: %+v\n", nums)
}

// 双指针解法
func func2(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
	fmt.Printf("func2: %+v\n", nums)
}

func func3(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	numsNew := make([]int, length)

	j := 0
	for _, val := range nums {
		if val != 0 {
			numsNew[j] = val
			j++
		}
	}
	fmt.Printf("func2: %+v\n", numsNew)
}

func main() {
	nums := []int{0, 1, 3, 0, 4}
	func1(nums)
	func2(nums)
	func3(nums)
}

/*============= 二刷 2-5 =============== */
func func52(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if j != i {
				nums[i] = 0
			}
			j++
		}
	}
	fmt.Printf("%+v", nums)
}

/*============= 二刷 3-5 =============== */
func func53(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if j != i {
				nums[i] == 0
			}
			j++
		}
	}
}
