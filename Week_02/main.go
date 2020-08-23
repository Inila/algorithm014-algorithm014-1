package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// heap 实现
	h := &IntHeap{1, 5, 2, 6, 9}
	heap.Init(h)
	heap.Push(h, 100)
	for h.Len() > 0 {
		fmt.Printf("%d\n ", heap.Pop(h))
	}

	// 滑动窗口最大值
	fmt.Printf("max: %d", findMaxByHeap([]int{1, 5, 2, 6, 9}))
}
