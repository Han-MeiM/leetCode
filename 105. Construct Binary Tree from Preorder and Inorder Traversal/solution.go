package main

import "github.com/Han-MeiM/leetCode/structure/TreeNode"

func buildTree(preorder []int, inorder []int) *TreeNode.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode.TreeNode{
			inorder[0],
			nil,
			nil,
		}
	}

	var left *TreeNode.TreeNode = nil
	var right *TreeNode.TreeNode = nil

	for k, v := range inorder {
		if v == preorder[0] {
			left = buildTree(preorder[1:k+1], inorder[0:k])
			right = buildTree(preorder[k+1:], inorder[k+1:])
			break
		}
	}

	root := &TreeNode.TreeNode{
		preorder[0],
		left,
		right,
	}

	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	TreeNode.Print(buildTree(preorder, inorder))
}
