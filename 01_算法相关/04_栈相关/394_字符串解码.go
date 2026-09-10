package main

import (
	"strings"
	"unicode"
)

// 394_字符串解码: 栈 + 临时缓存小对象
func decodeString(s string) string {
	// 创建一个对象存两个字段:
	type pair struct {
		ans string
		k   int
	}
	//创建一个栈
	stack := make([]pair, 0)
	// 初始化
	ans, k := "", 0
	// 遍历
	for _, c := range s {
		// 如果是字母,作为前面字符串
		if unicode.IsLetter(c) {
			ans += string(c)
		} else if unicode.IsDigit(c) {
			// 如果是数字,作为后面[]的倍数
			k = k*10 + int(c-'0')
		} else if c == '[' {
			// 如果是[,则先临时存起来
			stack = append(stack, pair{ans, k})
			// 重新下一轮
			ans, k = "", 0
		} else {
			// 如果是],则开始计算
			p := stack[len(stack)-1]
			// 栈顶元素弹出
			stack = stack[:len(stack)-1]
			ans = p.ans + strings.Repeat(ans, p.k)
		}
	}
	return ans
}
