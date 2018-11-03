package com.easy._237;

import com.structure.ListNode;

class Solution {
    public void deleteNode(ListNode node) {
        node.val = node.next.val;
        node.next = node.next.next;
    }

    public static void main(String args[]) {
        ListNode listNode = ListNode.createTestData("[4,5,1,9]");
        Solution solution = new Solution();
        solution.deleteNode(listNode.next);
        ListNode.print(listNode);
    }
}
