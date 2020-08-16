## 学习笔记

### 2020-08-10 23:35:00
下班较晚, 今天只看了数据、链表、跳表三种数据结构讲解视频。覃超老师讲的清晰、明了，突出重点。数据和链表比较熟悉，跳表之前不太了解。  
本次学习将跟随课程安排，使用覃超老师讲的算法学习方法，首先摒弃旧的习惯：

+ 不死磕问题，这是之前一直没坚持下来的原因，自己爱钻牛角尖
+ 五毒神掌，敢于方式，敢于自己硬背代码，要过最少5遍，而不是每次花费很长时间
    + 5-10 分钟读题和思考，如果没有思路，看题解，默写代码
    + 马上自己写，提交LeetCode，多种解法，体会优化
    + 24 小时之后，再重复做题
    + 一周后重复做题
    + 面试前一周恢复性训练
+ 不懒于看高手代码


### 数组 Arrary

+ 一段连续的内存，读取时间复杂度为O(1)
+ 插入和删除操作，时间复杂度为O(n), 因为需要挪动元素

### 链表 Linked List

+ 一段不连续的内存，读取需要从 head 开始遍历，时间复杂度为 O(n)
+ 插入和删除操作，直接需要next指针指向，时间复杂度为O(1)
+ 相较于数组，空间复杂度高，和数据互补，具体要看操作的应用场景
+ 不同的几种链表
    + 单向链表，tail 的 next指向null
    ```go
    // node 定义
    type SingleNode struct {
        Value int
        Next *SingleNode
    }

    // 单向链表
    type SingleList struct {
        Head *SingleNode
        Tail *SingleNode
        Size int
    }
    ```
    + 双向链表，除了 next 指向下一个 node，prev 指向前一个 node

    ```go
    // 双向链表 Node 定义
    type DoubleNode struct {
        Prev *DoubleNode
        Value int
        Next *DoubleNode
    }

    type LinkedList struct {
        Head *DoubleNode
        Tail *DoubleNode
        Size int
    }
    ```

    + 循环链表，tail 的 next 指向 head

### 跳表

+ 跳表只能用于元素有序的情况
+ 特点：原理简单、容易实现、方便扩展、效率更高
+ 用途：替代平衡树，如：Redis、LevelDB等
+ 实现方法：添加多级索引


### 解题相关

+ 最大的误区： 解题只解一遍
+ 优化思路： 升维， 空间换时间
+ 懵逼的时候怎么办？
    + 暴力解法？
    + 基本情况？
    + 找最近重复子问题