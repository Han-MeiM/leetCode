package com._013;

import java.util.Arrays;

public class Solution {
    public void reOrderArray(int[] array) {
        if (array == null || array.length == 0) return;
        int i = 0, j;
        while (i < array.length) {
            // 寻找偶数
            while (i < array.length && !this.iseven(array[i])) {
                i++;
            }

            // 寻找奇数
            j = i + 1;
            while (j < array.length && this.iseven(array[j])) {
                j++;
            }

            if (j < array.length) {
                // tmp 为奇数
                int tmp = array[j];
                for (int j2 = j - 1; j2 >= i; j2--) {
                    array[j2 + 1] = array[j2];
                }
                array[i] = tmp;
                i++;
            } else {// 查找失敗
                break;
            }
        }
    }

    private boolean iseven(int n) {
        return n % 2 == 0;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] array =  new int[]{1,2,3,4,5,6,7,8,9,10};
        solution.reOrderArray(array);
        System.out.println(Arrays.toString(array));
    }
}