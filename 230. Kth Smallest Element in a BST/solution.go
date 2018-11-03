package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	leftSize := calcTreeSize(root.Left)
	if k == leftSize + 1 {
		return root.Val
	} else if leftSize >= k {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k - leftSize - 1)
	}
}

// 获取左子树元素数目
func calcTreeSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + calcTreeSize(root.Left) + calcTreeSize(root.Right)
}

func main() {
	root := TreeNode{
		3,
		&TreeNode{
			1,
			nil,
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		&TreeNode{
			4,
			nil,
			nil,
		},
	}
	res := kthSmallest(&root, 1)
	fmt.Println(res)
}
