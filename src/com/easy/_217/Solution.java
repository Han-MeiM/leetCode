package com.easy._217;

import java.util.HashMap;

class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> temp = new HashMap<>();
        for (int value : nums) {
            if (temp.containsKey(value)) {
                return true;
            } else {
                temp.put(value, 1);
            }
        }
        return false;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        int[] nums = new int[]{1, 2, 3, 4};
        boolean res = solution.containsDuplicate(nums);
        System.out.println(res);
    }
}