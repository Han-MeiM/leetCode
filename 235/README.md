## 题名：[二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree)

根据LCA的定义和BST的特性

当root非空时可以区分为三种情况：

1）两个节点均在root的左子树，此时对root->left递归求解；

2）两个节点均在root的右子树，此时对root->right递归求解；

3）两个节点分别位于root的左右子树，此时LCA为root。