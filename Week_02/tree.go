package main

/*
	### 94. 二叉树的中序遍历
*/

// TreeNode: 二叉树节点定义
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 二叉树的遍历
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	st(root, &ret)
	return ret
}

func st(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	// 前序遍历
	/*
		*ret = append(*ret, root.Val)
		st(root.Left, ret)
		st(root.Right, ret)
	*/

	// 中序遍历
	st(root.Left, ret)
	*ret = append(*ret, root.Val)
	st(root.Right, ret)

	// 后续遍历
	/*
		st(root.Left, ret)
		st(root.Right, ret)
		*ret = append(*ret, root.Val)
	*/
}

// N叉数的遍历
type Node struct {
	Val       int
	Childrens []*Node
}

func postorder(root *Node) []int {
	rets := []int{}
	dfs(root, &rets)
	return rets
}

func dfs(root *Node, rets *[]int) {
	if root == nil {
		return
	}
	// 前序遍历
	*rets = append(*rets, root.Val)
	for _, children := range root.Childrens {
		dfs(children, rets)
	}
	// 后续遍历
	// *rets = append(*rets, root.Val)
}
