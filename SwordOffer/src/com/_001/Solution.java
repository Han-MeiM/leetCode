package com.array._001;

public class Solution {
    public boolean Find(int target, int[][] array) {
        if (array.length < 1 || array[0].length < 1) {
            return false;
        }
        int x_max = array[0].length;
        for (int[] y : array) {
            if (x_max == 0) return false;
            for (int x = 0; x < x_max; x++) {
                if (y[x] == target) return true;
                if (y[x] > target) {
                    x_max = x;
                    break;
                }
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
//        int[][] array = new int[][]{{1,2,8,9},{2,4,9,12},{4,7,10,13},{6,8,11,15}};
        int[][] array = new int[][]{{1, 2, 8, 9}, {4, 7, 10, 13},};
        boolean res = solution.Find(7, array);
        System.out.println(res);
    }
}
