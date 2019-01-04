package com._008;

public class Solution {
    public int JumpFloor(int target) {
        if (target < 2) {
            return 1;
        }
        int sum_1 = JumpFloor(target - 1);
        int sum_2 = JumpFloor(target - 2);
        return sum_1 + sum_2;
    }

    public int JumpFloorPlanB(int target) {
        if (target <= 0) {
            return -1;
        } else if (target == 1) {
            return 1;
        } else if (target == 2) {
            return 2;
        } else {
            return JumpFloorPlanB(target - 1) + JumpFloorPlanB(target - 2);
        }
    }

    public int JumpFloorPlanC(int target) {
        if (target == 0) return 0;
        if (target == 1) return 1;
        if (target == 2) return 2;
        int[] array = new int[target];
        array[0] = 1;
        array[1] = 2;
        for (int i = 2; i < target; i++) {
            array[i] = array[i - 1] + array[i - 2];
        }
        return array[target - 1];
    }

    public static void main(String[] args) {
        int target = 11;
        Solution solution = new Solution();
        int res = solution.JumpFloorPlanB(target);
        System.out.println(res);
    }
}
