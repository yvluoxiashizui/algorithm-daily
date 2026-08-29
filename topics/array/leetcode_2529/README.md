# 正整数和负整数的最大计数

- 力扣：https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/description/
- 日期：2026-08-29
- 难度：简单

## 题目
给你一个按 非递减顺序 排列的数组 nums ，返回正整数数目和负整数数目中的最大值。

换句话讲，如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值。
注意：0 既不是正整数也不是负整数

## 解题思路
sort.SearchInts它的返回值是第一个大于或等于目标值 x 的元素的索引位置。如果所有元素都小于 x，则返回切片的长度 len(a)

## 复杂度
- 时间：O(log n)
- 空间：O(1)