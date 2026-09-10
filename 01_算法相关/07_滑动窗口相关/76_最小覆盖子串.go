package main

/*
给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，返回空字符串 ""。
测试用例保证答案唯一。
示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。
提示：
m == s.length
n == t.length
1 <= m, n <= 105
s 和 t 由英文字母组成
解题思路:
滑动窗口 + 哈希表统计字符的次数
*/
// 最小覆盖子串: 哈希表统计元素出现的次数
func minWindow(s string, t string) string {
	// 定义一个哈希表的涵盖判断函数
	var isCovered func(m map[rune]int) bool
	isCovered = func(m map[rune]int) bool {
		for _, count := range m {
			if count > 0 {
				return false
			}
		}
		return true
	}
	// 首先创建一个哈希表统计t中的每个元素出现的次数
	m := make(map[rune]int)
	for _, c := range t {
		m[c]++
	}
	// 创建两个变量，存储窗口的左指针(设置为小于0,是为了当没有符合要求的时候,直接进行判断),右指针
	ansLeft, ansRight := -1, len(s)-1
	// 创建一个滑动窗口
	left := 0
	for right, c := range s {
		if _, ok := m[c]; ok {
			m[c]--
		}
		// 判断是否m中的各个元素的次数都是小于等于0
		for isCovered(m) {
			if right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}
			if _, exist := m[rune(s[left])]; exist {
				m[rune(s[left])]++
			}
			// 左指针右移动
			left++
		}
	}
	// 极端情况判断: 不存在涵盖的情况
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

// 76_最小覆盖子串: 定长数组统计次数 + 滑动窗口 + 双定长数组的涵盖判断
func minWindow2(s string, t string) string {
	// 定义一个定长数组涵盖判断
	var isCovered func(sm, tm [128]int) bool
	isCovered = func(sm, tm [128]int) bool {
		for i := 'A'; i <= 'Z'; i++ {
			if sm[i] < tm[i] {
				return false
			}
		}
		for i := 'a'; i <= 'z'; i++ {
			if sm[i] < tm[i] {
				return false
			}
		}
		return true
	}
	// 1. 创建一个定长数组来统计t的字符出现的次数: s和t都是由字母组成,所以用ascii码一定可以存下来
	tm := [128]int{}
	for _, c := range t {
		tm[c]++
	}
	// 2. 定义结果变量
	ansLeft, ansRight := -1, len(s)-1
	// 2. 定义滑动窗口的变量(这里具备单调性:因为左指针向右移动-count永远是-1,右指针向向右移动-count永远是+1)
	left, sm := 0, [128]int{}
	for right, c := range s {
		sm[c]++
		// 判断窗口是否涵盖
		for isCovered(sm, tm) {
			// 获取最小覆盖子串
			if right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}
			// 窗口左边缩容
			sm[s[left]]--
			// 窗口的左指针向右移动
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

// 76_最小覆盖子串: 定长数组统计字符出现次数 + 滑动窗口 + valid统计字符出现的种类
func minWindow3(s string, t string) string {
	// 创建一个定长数组tm来统计t的字符出现的次数；创建一个required常量来统计t的字符种类
	tm, required := [128]int{}, 0
	for _, c := range t {
		if tm[c] == 0 {
			required++
		}
		tm[c]++
	}
	// 定义结果变量
	ansLeft, ansRight := -1, len(s)-1
	// 创建滑动窗口变量
	left, valid := 0, 0
	for right, c := range s {
		tm[c]--
		if tm[c] == 0 {
			valid++
		}
		// 判断窗口是否涵盖
		for required == valid {
			if right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}
			// 获取即将缩容的字符
			lc := s[left]
			// 即将移除之前先判断是不是刚好满足,这个时候需要
			if tm[lc] == 0 {
				valid--
			}
			// 窗口缩容
			tm[lc]++
			left++
		}
	}
	// 边界处理
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}
