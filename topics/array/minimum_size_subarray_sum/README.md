# 长度最小的子数组

- 力扣：https://leetcode.cn/problems/minimum-size-subarray-sum/description/
- 日期：2026-08-24
- 难度：中等

## 题目
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 

## 解题思路
滑动窗口。右指针扩展加和，>= target 就收缩左指针记录最小长度。
出错点：ans 初始为 n+1 做哨兵，找不到就返回 0 易漏

## 复杂度
- 时间：O(n)
- 空间：O(1)