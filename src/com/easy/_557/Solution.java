package com.easy._557;

class Solution {
    public String reverseWords(String s) {
        String[] res = s.split(" ");
        StringBuilder str = new StringBuilder();
        for (int i = 0; i < res.length; i++) {
            str.append(Solution.reverseString(res[i]));
            if (i != res.length - 1) str.append(" ");
        }
        return str.toString();
    }

    private static String reverseString(String s) {
        int length = s.length();
        char[] res = s.toCharArray();
        for (int i = 0; i < length / 2; i++) {
            char t = res[i];
            res[i] = res[length - i -1];
            res[length - i - 1] = t;
        }
        return new String(res);
    }

    public static void main(String args[]) {
        String str = "Let's take LeetCode contest";
        Solution solution = new Solution();
        String res = solution.reverseWords(str);
        System.out.print(res);
    }
}
