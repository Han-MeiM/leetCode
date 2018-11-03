package com.easy._344;

class Solution {
    public String reverseString(String s) {
        int length = s.length();
        char[] res = s.toCharArray();
        for (int i = 0; i < length / 2; i++) {
            char t = res[i];
            res[i] = res[length - i -1];
            res[length - i - 1] = t;
        }
        return new String(res);
    }

    public static void main(String args[]){
        String str = "hello";
        Solution solution = new Solution();
        String res = solution.reverseString(str);
        System.out.print(res);
    }
}
