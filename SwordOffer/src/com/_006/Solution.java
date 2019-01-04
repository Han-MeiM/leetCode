package com._006;

public class Solution {
    public int minNumberInRotateArray(int[] array) {
        if (array.length <= 0) {
            return 0;
        }
        int low = 0;
        int high = array.length - 1;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (array[mid] > array[high]) {
                low = mid + 1;
            } else if (array[mid] == array[high]) {
                high = high - 1;
            } else {
                high = mid;
            }
        }
        return array[low];
    }

    // O(n)
    public int minNumberInRotateArrayPlanB(int[] array) {
        if (array.length <= 0) {
            return 0;
        }
        int befor = array[0];
        for (int item : array) {
            if (item < befor) {
                return item;
            }
            befor = item;
        }
        return array[0];
    }

    public static void main(String[] args) {
        int[] array = {3, 4, 5, 1, 2};
        Solution solution = new Solution();
        int res = solution.minNumberInRotateArray(array);
        System.out.println(res);
    }
}
