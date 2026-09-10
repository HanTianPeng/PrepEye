package main

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/
// 438_找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	// 创建一个固定长度的数组,来统计异位词的每个字母的出现次数
	// 因为字母最多只有26个
	countList := make([]int, 26)
	for _, c := range p {
		// 利用差值计算;a为最小
		countList[c-'a']++
	}
	// 创建一个动态的窗口;
	left := 0
	for right, c := range s {
		// 利用差值计算当前字母的数字:
		c = c - 'a'
		// countList是目标的数组;利用减法进行计算;
		// 正常减法过后每个元素应该count数量为0
		countList[c]--
		// 如果减法过后,该数字小于0;存在两种情况
		// 第一种情况,该字母不在异位词的里面 -> left指针右移动一个位置其实还是浪费了些时间复杂度
		// 第二种情况,该字母在异位词里面但是太多了 -> left指针右移动一个位置
		for countList[c] < 0 {
			// 这个时候就需要把窗口右移动一个位置;
			// 将原来的目标数组进行恢复
			countList[s[left]-'a']++
			left++
		}
		if right-left+1 == len(p) {
			ans = append(ans, left)
		}
	}
	return ans
}

func findAnagrams2(s string, p string) []int {
	ans := make([]int, 0)
	// 统计异位词的每个字母出现的次数
	countList := [26]int{}
	for _, c := range p {
		countList[c-'a']++
	}
	// 创建一个固定长度的滑动窗口
	dynamicList := [26]int{}
	for right, c := range s {
		dynamicList[c-'a']++
		left := right - len(p) + 1
		// 字母的数量不满足这个异位词的数量
		if left < 0 {
			continue
		}
		// 判断固定长度构成的字符串是否符合异位词
		if dynamicList == countList {
			ans = append(ans, left)
		}
		// 将滑动窗口往右边移动一个位置
		dynamicList[s[left]-'a']--
	}
	return ans
}
