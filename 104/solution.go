package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left+1
	} else {
		return right+1
	}
}

///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	cur := 0
//	res := recursion(root, cur)
//	return res
//}
//
//func recursion(node *TreeNode, cur int) int {
//	cur++
//	left := cur
//	right := cur
//	if node.Left != nil {
//		left = recursion(node.Left, cur)
//	}
//	if node.Right != nil {
//		right = recursion(node.Right, cur)
//	}
//
//	if left > right {
//		return left
//	} else {
//		return right
//	}
//}

func main() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 7,
				Left: nil,
				Right: nil,
			},
		},
	}
	res := maxDepth(&root)
	fmt.Println(res)
}
