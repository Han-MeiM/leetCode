package com.easy._021;

import com.structure.ListNode;

class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode res;
        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l1;
        }

        if (l1.val < l2.val) {
            ListNode temp = mergeTwoLists(l1.next, l2);
            res = l1;
            res.next = temp;
        } else {
            ListNode temp = mergeTwoLists(l1, l2.next);
            res = l2;
            res.next = temp;
        }
        return res;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        ListNode res = solution.mergeTwoLists(ListNode.createTestData("[1,2,4]"), ListNode.createTestData("[1,3,4]"));
        ListNode.print(res);
    }
}
