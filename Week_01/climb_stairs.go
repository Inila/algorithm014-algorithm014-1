package main

import "fmt"

/*
	70. 爬楼梯

	思路：懵逼中。。。

	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	注意：给定 n 是一个正整数。

	我们用 f(x)f(x) 表示爬到第 xx 级台阶的方案数，考虑最后一步可能跨了一级台阶，也可能跨了两级台阶，所以我们可以列出如下式子：
		f(x) = f(x - 1) + f(x - 2)
		f(x)=f(x−1)+f(x−2)

*/

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	fmt.Printf("%d\n", r)
	return r
}

func main() {
	climbStairs(30)
}
