package com._010;

public class Solution {
    public int RectCover(int target) {
        com._008.Solution solution = new com._008.Solution();
        return solution.JumpFloorPlanC(target);
    }

    public static void main(String[] args) {
        int target = 11;
        Solution solution = new Solution();
        int res = solution.RectCover(target);
        System.out.println(res);
    }
}
