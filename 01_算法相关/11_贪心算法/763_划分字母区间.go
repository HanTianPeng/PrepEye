package main

/*
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
返回一个表示每个字符串片段的长度的列表。
示例 1：
输入：s = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。
示例 2：
输入：s = "eccbbbbdec"
输出：[10]
提示：
1 <= s.length <= 500
s 仅由小写英文字母组成
*/
// 划分字母区间: 统计字母出现次数 + 贪心算法 + 双指针
func partitionLabels(s string) []int {
	ans := make([]int, 0)
	marks := make([]int, 26)
	// 计算字符出现的最大索引
	for i := 0; i < len(s); i++ {
		marks[s[i]-'a'] = i
	}
	// 遍历字符串
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		mark := marks[s[i]-'a']
		right = max(mark, right)
		if i == right {
			ans = append(ans, i-left+1)
			left = i + 1
		}
	}
	return ans
}

func partitionLabels2(s string) []int {
	// 找到每个字符的最远索引
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	maxIdx := 0
	left := 0
	ans := make([]int, 0)
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 找到该字符串的最远索引
		curIdx := m[s[i]]
		if curIdx > maxIdx {
			maxIdx = curIdx
		}
		if i == maxIdx {
			ans = append(ans, i-left+1)
			left = i + 1
		}
	}
	return ans
}
