package com.easy._206;

import com.structure.ListNode;

public class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode res = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return res;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        ListNode head = ListNode.createTestData("[1,2,3,4,5]");
        ListNode res = solution.reverseList(head);
        ListNode.print(res);
    }
}
