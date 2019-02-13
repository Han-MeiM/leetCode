package main

import (
	"github.com/Han-MeiM/leetCode/structure/TreeNode"
)

func lowestCommonAncestor(root, p, q *TreeNode.TreeNode) *TreeNode.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 判断左边或右边是否存在 p 或 q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	} else {
		return left
	}
}

func main() {
	root := TreeNode.CreateTestData("[3,5,1,6,2,0,8,nil,nil,7,4]")
	res := lowestCommonAncestor(root, root.Left, root.Right)
	TreeNode.Print(res)
}
