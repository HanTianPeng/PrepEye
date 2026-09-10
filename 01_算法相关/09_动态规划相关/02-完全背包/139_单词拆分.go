package main

/*
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
说明：
拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
注意你可以重复使用字典中的单词。
示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
#
解题思路:
wordDict为金钱、s的钱包
1. 定义dp[j]的含义及下标的含义:
	dp[j]: 用wordDict的单词,装满长度为j的字符串,有多少种方法
2. 动态转换方程式:
	dp[j] += dp[j - len(wordDict[i])]
3. 初始化:
	dp[0] = 1
4. 遍历顺序及判断条件:
	完全背包问题,组合情况,所以顺序不重要;
	当背包的长度为j的时候, 单词的长度一定要小于背包的长度j,不然这个背包也装不下单词wordDict[i]

如果dp[j]=true;代表j个字符构成的字符串可以拆分为一个或多个在字典中出现的单词
如果dp[i]=true,则证明在[i,j]的区间的字符构成的单词在字典中出现
所以如果[i,j]的区间的字符构成的单词在字典中出现并且dp[i]=true，则证明dp[j]=true
*/

func wordBreak(s string, wordDict []string) bool {
	// 创建完全背包动态规划的一维数组
	packageWeight := len(s)
	dp := make([]int, packageWeight+1)
	// 初始化
	dp[0] = 1
	// 动态规划遍历: 先遍历背包,再遍历商品
	for j := 1; j <= packageWeight; j++ {
		for i := 0; i < len(wordDict); i++ {
			// 这里为什么可以解决单词拼接顺序问题,因为j是从1开始,无论子串是在前面还是中间或是后面,都会从头开始截取匹配
			if len(wordDict[i]) <= j && wordDict[i] == s[j-len(wordDict[i]):j] {
				dp[j] += dp[j-len(wordDict[i])]
			}
		}
	}
	return dp[packageWeight] > 0
}
