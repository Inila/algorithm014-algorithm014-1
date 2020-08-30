package main

import "fmt"

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/

/*
	思考思路:
	1. n代表生成括号的对数，那么实际上共有 2*n 个格子可以放括号
	2. 通过 generateParenthesis1 列出了所有的可能性，但是此时没有办法保证括号的合法性 (到这一次竟然自己递归一次就写出来了。。。)
	3. 关于生成的字符合法性校验:
		a. 对所有的可能性结果进行合法性校验，使用stack，显然每个结果都要再校验一遍，时间复杂度比较高
		b. 直接在递归中进行合法性校验

	// 如果在递归中进行合法性校验呢？
	1. 左括号 或者 右括号不能超过 n 次
	2. 写入右括号时, 左括号的次数需要大于右括号的次数
	满足上面这两个条件就能排除所有非合法的字符串嘛。。其实我也不确定，这里先硬记住吧。。。
*/
func generateParenthesis1(n int) []string {
	rets := make([]string, 0)
	//
	dfs1(0, n*2, "", &rets)
	return rets
}

func dfs1(level, maxLevel int, s string, rets *[]string) {
	// level 当前的位置
	// maxLevel 最大的位置，即: n*2
	// rets 保存最终结果
	// currentS 当前括号的字符串
	// s 即放入的括号, 可以是左括号或者右括号

	// 退出条件
	if level > maxLevel {
		*rets = append(*rets, s)
		return
	}
	// 本次递归需要执行的程度代码， 向 rets 中放入 左括号 或者 右括号
	dfs1(level+1, maxLevel, s+"(", rets)
	dfs1(level+1, maxLevel, s+")", rets)
}

func main() {
	// 递归遍历生成所有的可能性
	// rets1 := generateParenthesis1(3)
	// fmt.Printf("rets1: %+v\n", rets1)

	// 优先解法
	rets := generateParenthesis(5)
	fmt.Printf("rets: %+v\n", rets)
}

func generateParenthesis(n int) []string {
	rets := new([]string)
	//
	dfs(0, 0, n, "", rets)
	return *rets
}

func dfs(leftCnt, rightCnt int, max int, s string, rets *[]string) {
	// level 遍历层数
	// n 括号对的个数，也为 leftCnt 和 rightCnt 的最大次数
	// currentS 当前生成的括号字符串
	// s 需要添加左括号或者右括号
	// rets 最终生成的结果

	// 这里的退出条件需要记住。。。
	if leftCnt == rightCnt && leftCnt == max {
		*rets = append(*rets, s)
		return
	}

	// 左括号的次数还没有达到最大次数
	if leftCnt < max {
		dfs(max, leftCnt+1, rightCnt, s+"(", rets)
	}

	if rightCnt < leftCnt {
		dfs(max, leftCnt, rightCnt+1, s+")", rets)
	}
}
