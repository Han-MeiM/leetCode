package com._012;

public class Solution {
    public double Power(double base, int exponent) {
        if (exponent == 0) {
            return 1;
        } else if (exponent < 0) {
            if (base == 0) {
                throw new RuntimeException("分母不能为0");
            }
            base = 1 / base;
            exponent = -exponent;
        }
        double res = base;
        while (exponent > 1) {
            res = res * base;
            exponent--;
        }
        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        double res = solution.Power(2, -3);
        System.out.println(res);
    }
}
