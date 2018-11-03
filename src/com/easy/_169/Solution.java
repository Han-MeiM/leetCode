package com.easy._169;

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int majorityElement(int[] nums) {
        int n = nums.length, res = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < n; i++) {
            int count;
            if (map.get(nums[i]) == null) {
                count = 1;
            } else {
                count = map.get(nums[i]) + 1;
            }
            if (count > n / 2) {
                res = nums[i];
                break;
            }
            map.put(nums[i], count);
        }
        System.out.print(map.toString());
        return res;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        int res = solution.majorityElement(new int[]{2,2,1,1,1,2,2});
        System.out.print(res);
    }
}
