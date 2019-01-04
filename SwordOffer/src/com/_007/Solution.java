package com._007;

public class Solution {
    public int Fibonacci(int n) {
        int f = 0, g = 1;
        // f(n) = f(n+1) - f(n-1)
        while (n > 0) {
            g += f;
            f = g - f;
            n--;
        }
        return f;
    }

    public static void main(String[] args) {
        int n = 6;
        Solution solution = new Solution();
        int res = solution.Fibonacci(n);
        System.out.println(res);
    }
}