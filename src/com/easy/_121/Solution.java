package com.easy._121;

class Solution {
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length < 1) {
            return 0;
        }
        int min = prices[0];
        int max = 0;
        for (int value : prices) {
            if (min > value) {
                min = value;
            } else {
                if (max < value - min) {
                    max = value - min;
                }
            }

        }
        return max;
    }

    public static void main(String args[]) {
        int[] prices = new int[]{7, 1, 5, 3, 6, 4};
        Solution solution = new Solution();
        int res = solution.maxProfit(prices);
        System.out.println(res);
    }
}