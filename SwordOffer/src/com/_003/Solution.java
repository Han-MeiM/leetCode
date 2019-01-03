package com.list._003;

import com.structure.ListNode;
import java.util.ArrayList;

public class Solution {
    public ArrayList<Integer> printListFromTailToHead(ListNode ListNode) {
        if (ListNode == null) {
            return new ArrayList<>();
        }
        ArrayList<Integer> list = this.printListFromTailToHead(ListNode.next);
        list.add(ListNode.val);
        return list;
    }

    public static void main(String[] args) {
        ListNode list = ListNode.createTestData("[1,2,3,4,5,6,7,8]");
        Solution solution = new Solution();
        ArrayList<Integer> res = solution.printListFromTailToHead(list);
        System.out.println(res);
    }
}
