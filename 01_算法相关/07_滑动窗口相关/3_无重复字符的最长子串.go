package main

/*
滑动窗口 = 同向双指针 + 窗口状态
同向双指针:
	left: 负责收缩窗口
	right: 负责扩张窗口
	两个指针都只往右走
窗口状态: 解决当前窗口是否满足条件
	字符频次
	哈希集合
	当前区间和
	单调队列
窗口条件: 必须单调性
写代码前需要考虑三个问题:
1. 窗口里要维护什么状态?
2. 什么情况下需要移动left?
3. 答案在缩窗前更新还是缩窗后更新?
*/
// 3_无重复字符的最长子串: 滑动窗口 + 字符串是否重复 + 频次数组写法
func lengthOfLongestSubstring(s string) int {
	// 创建一个数组，用来判断字符串是否重复
	cnt := [128]int{}
	// 定义滑动窗口变量
	left, ans := 0, 0
	// 滑动窗口遍历
	for right, c := range s {
		cnt[c]++
		// 窗口的状态: 字符串重复
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// 3_无重复字符的最长子串: 滑动窗口 + 字符串是否重复 + lastIndex写法
func lengthOfLongestSubstring2(s string) int {
	// 创建一个数组，用来判断字符串是否重复
	cnt := [128]int{}
	// 定义滑动窗口变量
	left, ans := 0, 0
	// 滑动窗口遍历
	for right, c := range s {
		cnt[c] = right
		// 窗口的状态:
		if cnt[c] > left {
			left = cnt[c] + 1
		}

		ans = max(ans, right-left+1)
	}
	return ans
}

// 3_无重复字符的最长子串: 滑动窗口 + 字符串重复的索引位置
func lengthOfLongestSubstring3(s string) int {
	// 创建一个窗口
	list := make([]int32, 0)

	// 遍历字符串
	ans := 0
	for _, str := range s {
		// 查询窗口中是否存在该字符并返回索引
		idx := findIdxByStr(list, str)
		// 如果不重复，则直接添加到滑动窗口中
		if idx == -1 {
			list = append(list, str)
			ans = max(ans, len(list))
		} else {
			// 计算之前的最大值
			ans = max(ans, len(list))
			// 将重复元素添加到窗口中
			list = append(list, str)
			// 将滑动窗口进行移动
			list = list[idx+1:]
		}
	}
	return ans
}

func findIdxByStr(list []int32, str int32) int {
	idx := -1
	for i, s := range list {
		if str == s {
			idx = i
			break
		}
	}
	return idx
}
