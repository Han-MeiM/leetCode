package com._002;

public class Solution {
    public String replaceSpace(StringBuffer str) {
        int spaceCount = 0;
        // 计算空格数
        for (int i = 0; i < str.length(); i++) {
            if (str.charAt(i) == ' ') {
                spaceCount++;
            }
        }
        int strOldLength = str.length(); // str 的旧长度
        int strNewLength = strOldLength + spaceCount * 2; // str 的新长度
        str.setLength(strNewLength); // 给 str 扩容
        for (int j = strOldLength - 1; j >= 0; j--) {
            if (str.charAt(j) != ' ') {
                str.setCharAt(j + spaceCount * 2, str.charAt(j));
            } else {
                spaceCount--;
                str.setCharAt(j + spaceCount * 2, '%');
                str.setCharAt(j + spaceCount * 2 + 1, '2');
                str.setCharAt(j + spaceCount * 2 + 2, '0');
            }
        }
        return str.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        StringBuffer str = new StringBuffer("asd asd asd");
        System.out.println(solution.replaceSpace(str));
    }
}
