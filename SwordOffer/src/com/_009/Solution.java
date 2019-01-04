package com._009;

public class Solution {
    public int JumpFloorII(int target) {
        if (target <= 0) {
            return -1;
        } else if (target == 1) {
            return 1;
        } else {
            return 2 * JumpFloorII(target - 1);
        }
    }

    public static void main(String[] args) {
        int target = 11;
        Solution solution = new Solution();
        int res = solution.JumpFloorII(target);
        System.out.println(res);
    }
}
