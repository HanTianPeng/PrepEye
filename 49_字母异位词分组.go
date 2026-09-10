package main

import (
	"maps"
	"slices"
)

// 1. 字符串随机组合排序最终的结果永远一致 + slices.Sort方法 + maps.Values方法 + slices.Collect方法
func groupAnagrams(strs []string) [][]string {
	// 定义一个map数据结构,key是string,value是[]string
	mp := make(map[string][]string)

	// 对字符串进行遍历
	for _, str := range strs {
		// 将字符串转换成[]byte;并进行排序
		// 这里面有一个非常大的时间复杂度的开销;每次都要对str进行排序
		temp := []byte(str)
		slices.Sort(temp)
		mp[string(temp)] = append(mp[string(temp)], str)
	}
	return slices.Collect(maps.Values(mp))
}

// 2. 字符串随机组合排序最终每个字母出现的次数永远相等 + map的key为[26]int + [26]int初始化[26]int{}
func groupAnagrams2(strs []string) [][]string {
	// 定义一个map的数据结构,key是[]int,value是[]string
	mp := make(map[[26]int][]string)
	// 对字符串进行遍历
	for _, str := range strs {
		key := [26]int{}
		for _, s := range str {
			key[s-'a']++
		}
		// 固定长度的数组可以作为map的key字段;这就是比较巧妙的编写方法
		mp[key] = append(mp[key], str)
	}
	return slices.Collect(maps.Values(mp))
}
