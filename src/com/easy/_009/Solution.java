package com.easy._009;

class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }

        int temp = 0;
        while (temp < x) {
            temp = x % 10 + temp * 10;
            x /= 10;
        }
        return temp == x || x == temp / 10;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        int test = 121;
        boolean res = solution.isPalindrome(test);
        System.out.print(res);
    }
}
