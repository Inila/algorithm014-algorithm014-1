package main

/*
	官方库中 container/heap 实现了堆的接口操作，接口定义如下:

	type Interface interface {
    	sort.Interface
    	Push(x interface{}) // 向末尾添加元素
    	Pop() interface{}   // 从末尾删除元素
	}
*/

/*
	Int 最小堆实现
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Less 方式这样定义就是大根堆
// func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

//Push 向heap中添加元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

//Pop 从heap中取出元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
