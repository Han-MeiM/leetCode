package com.easy._004;

public class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int lengthA = nums1.length;
        int lengthB = nums2.length;
        int len = lengthA + lengthB;
        if (len % 2 == 0) {
            return (findKth(nums1, nums2, 0, 0, len / 2) + findKth(nums1, nums2, 0, 0, len / 2 + 1)) / 2.0;
        } else {
            return findKth(nums1, nums2, 0, 0, len / 2 + 1);
        }
    }

    private static double findKth(int[] nums1, int[]nums2, int starta, int startb, int k){
        int len1 = nums1.length;
        int len2 = nums2.length;

        if(starta >= len1){
            return nums2[startb + k - 1];
        }
        if(startb >= len2){
            return nums1[starta + k - 1];
        }
        if(k == 1){
            return Math.min(nums1[starta], nums2[startb]);
        }

        int mid = k/2 - 1;
        int keypoint1 = starta + mid >= len1? Integer.MAX_VALUE : nums1[starta + mid];
        int keypoint2 = startb + mid >= len2? Integer.MAX_VALUE : nums2[startb + mid];

        if(keypoint1 > keypoint2){
            return findKth(nums1, nums2, starta, startb + k/2, k - k/2);
        } else {
            return findKth(nums1, nums2, starta + k/2, startb, k - k/2);
        }
    }

    public static void main(String[] args) {
        int[] nums1 = new int[]{1, 2};
        int[] nums2 = new int[]{3, 4};
        Solution solution = new Solution();
        double res = solution.findMedianSortedArrays(nums1, nums2);
        System.out.println(res);
    }
}
