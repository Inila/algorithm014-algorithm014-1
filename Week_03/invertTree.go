package main

/*
	翻转一棵二叉树。

	输入：
	      4
   		/   \
  	   2     7
 	  / \   / \
	 1   3 6   9

	 输出：
	      4
   		/   \
  	   7     2
 	  / \   / \
	9   6  3   1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归牛B...
func invertTree(root *TreeNode) *TreeNode {
	// 前序遍历
	if root == nil {
		return nil
	}

	right := invertTree(root.Right)
	left := invertTree(root.Left)

	root.Right = left
	root.Left = right

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	//
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
