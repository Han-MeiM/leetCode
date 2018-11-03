package com.easy._005;

public class Solution {
    public String longestPalindrome(String s) {
        int n = s.length(), start = 0, maxLen = 1;
        boolean[][] dp = new boolean[n][n];

        for (int i = 0; i < n; i++) {
            dp[i][i] = true;
            if (i + 1 < n && s.charAt(i) == s.charAt(i + 1)) {
                dp[i][i + 1] = true;
                start = i;
                maxLen = 2;
            }
        }

        for (int i = n - 1; i >= 0; i--) {
            for (int j = i + 2; j < n; j++) {
                if (dp[i + 1][j - 1] && s.charAt(i) == s.charAt(j)) {
                    dp[i][j] = true;
                    if (j - i >= maxLen) {
                        start = i;
                        maxLen = j - i + 1;
                    }
                }
            }
        }
        System.out.println(start);
        System.out.println(maxLen);
        return s.substring(start, maxLen + 1);
    }

    public static void main(String[] args) {
        String str = "babaddab";
        Solution solution = new Solution();
        String res = solution.longestPalindrome(str);
        System.out.println(res);
    }
}
