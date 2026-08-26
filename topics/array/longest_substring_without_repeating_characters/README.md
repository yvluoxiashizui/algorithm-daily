# 长度最小的子数组

- 力扣：https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
- 日期：2026-08-26
- 难度：中等

## 题目
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

## 解题思路
滑动窗口+哈希表。滑动窗口同前题，用num储存出现次数，num>1说明出现重复，删去

## 复杂度
- 时间：O(n)
- 空间：O(1)