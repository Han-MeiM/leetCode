package com.easy._235;

import com.structure.TreeNode;

class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        TreeNode res;
        if ((p.val >= root.val && q.val <= root.val) || (p.val <= root.val && q.val >= root.val)) {
            res = root;
        } else if (p.val < root.val) {
            res = lowestCommonAncestor(root.left, p, q);
        } else {
            res = lowestCommonAncestor(root.right, p, q);
        }
        return res;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        TreeNode root = TreeNode.createTestData("[6,2,8,0,4,7,9,null,null,3,5]");
        TreeNode res = solution.lowestCommonAncestor(root, root.left, root.right);
        TreeNode.print(res);
    }
}
