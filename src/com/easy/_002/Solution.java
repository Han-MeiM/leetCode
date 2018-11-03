package com.easy._002;

import com.structure.ListNode;

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode res = new ListNode(0);
        ListNode temp = res;
        int sum = 0;
        while (l1 != null || l2 != null) {
            if (l1 != null) {
                sum += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                sum += l2.val;
                l2 = l2.next;
            }
            temp.next = new ListNode(sum % 10);
            temp = temp.next;
            sum = sum / 10;
        }
        if (sum != 0) {
            temp.next = new ListNode(sum);
        }
        return res.next;
    }

    public static void main(String[] args) {
        ListNode l1 = ListNode.createTestData("[2,4,3]");
        ListNode l2 = ListNode.createTestData("[5,6,4]");
        Solution solution = new Solution();
        ListNode res = solution.addTwoNumbers(l1, l2);
        System.out.println(res);
    }
}
