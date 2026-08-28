# 在排序数组中查找元素的第一个和最后一个位置

- 力扣：https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
- 日期：2026-08-26
- 难度：中等

## 题目
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
## 解题思路
二分，注意寻找右边界是寻找目标值加一的边界，然后减一得到右边界。
几种情况还没有完全掌握，标记

## 复杂度
- 时间：O(n)
- 空间：O(1)