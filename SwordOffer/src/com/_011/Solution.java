package com._011;

public class Solution {
    public int NumberOf1(int n) {
        int count = 0;
        while (n != 0) {
            ++count;
            n = (n - 1) & n;
        }
        return count;
    }

    public int NumberOf1PlanB(int n) {
        int count = 0;
        int m = 1;
        while (m != 0) {
            if ((n & m) != 0) {
                count++;
            }
            m = m << 1;
        }
        return count;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int res = solution.NumberOf1PlanB(20);
        System.out.println(res);
    }
}
