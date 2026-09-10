package main

import (
	"strconv"
	"strings"
)

// 93_复原IP地址:
func restoreIpAddresses(s string) []string {
	path, ans := make([]string, 0), make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == 4 {
			// 如果已经切了4段，还没有包括所有的字符，证明是无意义切割
			if start != len(s) {
				return
			}
			str := strings.Join(path, ".")
			ans = append(ans, str)
			return
		}
		// 裁剪: 每次切割位置必须要留还未切割次数的量
		for i := start; i < len(s)-(4-len(path))+1; i++ {
			// 含有前导0不符合ip规范: s[start]而不是s[i]
			if i != start && s[start] == '0' {
				break
			}
			x := s[start : i+1]
			num, _ := strconv.Atoi(x)
			// 判断是否符合ip规范
			if num >= 0 && num <= 255 {
				path = append(path, x)
				dfs(i + 1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	dfs(0)
	return ans
}
