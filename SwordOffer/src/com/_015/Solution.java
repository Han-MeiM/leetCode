package com._015;

import com.structure.ListNode;

public class Solution {
    public ListNode ReverseList(ListNode head) {
        if (head == null || head.next == null) return head;

        // 原链表 尾节点，新链表 头节点
        ListNode res = ReverseList(head.next);
        head.next.next = head;
        head.next = null;
        return res;
    }

    public static void main(String[] args) {
        ListNode head = ListNode.createTestData("[1,2,3,4,5,6,7]");
        Solution solution = new Solution();
        ListNode res = solution.ReverseList(head);
        ListNode.print(res);
    }
}