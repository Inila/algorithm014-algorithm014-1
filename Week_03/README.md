### 学习笔记

### go 递归模板代码

```text
func recur(level int, param int) { 
  // terminator 退出条件
  if (level > MAX_LEVEL) { 
    // process result 
    return; 
  }

  // process current logic 
  process(level, param); 
  
  // drill down 
  recur( level: level + 1, newParam); 
  
  // restore current status 
}
```
### 递归的实现、特性以及思维要点
这个总结的是真好，套路好使，括号生成问题，中午看的视频，下午忙别的去，晚上看这问题慢慢回忆，竟然很神奇的自己写出来了，而且接进最优答案，这是目前为止第一道题这么写出来的，其他题目第一次解答全都靠题解~


### 分治和回溯
实话实说，这两节视频看完了觉得有些懵逼。。。也不是没听懂，而是觉得怎么来应用这些思想，完全没思路。。我还是先多刷刷题，慢慢领悟吧~ 等着后续再看几遍~


### 相关习题
+ [爬楼梯问题](climbStairs.go)
+ [括号生成](generateParenthesis.go)
+ [反转二叉树](invertTree.go)