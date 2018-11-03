package com.easy._292;

class Solution {
    public boolean canWinNim(int n) {
        if (n % 4 == 0) {
            return false;
        }
        return true;
    }

    public static void main(String args[]) {
        Solution solution = new Solution();
        boolean res = solution.canWinNim(5);
        System.out.print(res);
    }
}
