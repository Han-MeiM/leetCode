package com.easy._104;

import com.structure.TreeNode;

class Solution {
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int left = maxDepth(root.left);
        int right = maxDepth(root.right);
        if (left < right) {
            return right + 1;
        }
        return left + 1;
    }

    public static void main(String[] args) {
        TreeNode root = TreeNode.createTestData("[3,9,20,null,null,15,7]");
        Solution solution = new Solution();
        int res = solution.maxDepth(root);
        System.out.print(res);
    }
}
