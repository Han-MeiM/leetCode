package com.easy._238;

import java.util.Arrays;

class Solution {
    public int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int right = 1;
        int[] res = new int[n];
        res[0] = 1;
        for (int i = 1; i < n; i++) {
            res[i] = res[i - 1] * nums[i - 1];
        }

        for (int i = n - 2; i >= 0; i--) {
            right *= nums[i + 1];
            res[i] *= right;
        }

        return res;
    }

    public static void main(String args[]) {
        int[] nums = new int[]{1,2,3,4};
        Solution solution = new Solution();
        int[] res = solution.productExceptSelf(nums);
        System.out.println(Arrays.toString(res));
    }
}
