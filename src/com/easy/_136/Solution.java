package com.easy._136;

class Solution {
    public int singleNumber(int[] nums) {
        int res = 0;
        for (int i = 0; i < nums.length; i++) {
            res = res ^ nums[i];
        }
        return res;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        int res = solution.singleNumber(new int[]{4,1,2,1,2});
        System.out.print(res);
    }
}
