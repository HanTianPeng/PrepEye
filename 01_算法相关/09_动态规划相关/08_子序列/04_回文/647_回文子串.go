package main

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
回文字符串 是正着读和倒过来读一样的字符串。
子字符串 是字符串中的由连续字符组成的一个序列。
示例 1：
输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：
输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
提示：
1 <= s.length <= 1000
s 由小写英文字母组成
解题思路:
1. 确定dp[i][j]的定义及下标的含义
	dp[i][j]: 表示区间范围[i,j]的子串是否是回文子串,如果dp[i]][j]为true,否则为false
2. 递归公式的推导
	s[i]与s[j]不相等: false
	s[i]与s[j]相等:
		1. i == j: true
		2. j - i = 1: true
		3. j - i > 1: dp[i+1][j-1]是否为true
3. 初始化
4. 递归遍历顺序
	i=1		j=1		j=2		j=j-1			j=j
	i=2										dp[i][j]
	i=i
	i=i+1					dp[i+1][j-1]
如果这矩阵是从上到下，从左到右遍历，那么会用到没有计算过的dp[i + 1][j - 1]，
也就是根据不确定是不是回文的区间[i+1,j-1]，来判断了[i,j]是不是回文，那结果一定是不对的。
所以一定要从下到上，从左到右遍历，这样保证dp[i + 1][j - 1]都是经过计算的。
*/
// 647_回文子串: 动态规划二维数组
func countSubstrings(s string) int {
	l := len(s)
	// 创建动态规划二维数组变量dp[i][j]: 表示区间范围[i,j]的子串是否是回文子串,如果dp[i]][j]为true,否则为false
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
	}
	// 创建一个结果变量
	ans := 0
	// 动态规划遍历
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				ans++
				dp[i][j] = true
			}
		}
	}
	return ans
}

// 回文子串: 双指针
func countSubstrings2(s string) int {
	res := 0
	l := len(s)
	// 定义一个函数计算以某个点或某两个点为中心构成的回文字符串的数量
	var countPalindromic func(i, j int) int
	countPalindromic = func(i, j int) int {
		ans := 0
		for i >= 0 && j < l && s[i] == s[j] {
			i--
			j++
			ans++
		}
		return ans
	}
	for i := 0; i < l; i++ {
		res += countPalindromic(i, i)
		res += countPalindromic(i, i+1)
	}
	return res
}
