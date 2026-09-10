package main

// 46_全排列: nums为不重复的元素
func permute(nums []int) [][]int {
	l := len(nums)
	path, used, ans := make([]int, 0, l), make([]bool, l), make([][]int, 0)
	var dfs func()
	dfs = func() {
		if len(path) == l {
			temp := make([]int, l)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < l; i++ {
			// 已经取过的元素不再重复提取
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return ans
}
