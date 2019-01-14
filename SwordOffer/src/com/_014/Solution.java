package com._014;

import com.structure.ListNode;

public class Solution {
    public ListNode FindKthToTail(ListNode head, int k) {
        if (head == null || k == 0) return null;

        ListNode fast = head, slow = head;
        while (k > 1) {
            if (fast.next == null) return null;
            fast = fast.next;
            k--;
        }
        while (fast.next != null) {
            fast = fast.next;
            slow = slow.next;
        }
        return slow;
    }

    public static void main(String[] args) {
        ListNode head = ListNode.createTestData("[1,2,3,4,5,6,7]");
        Solution solution = new Solution();
        ListNode res = solution.FindKthToTail(head, 4);
        ListNode.print(res);
    }
}
