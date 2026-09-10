package main

import "strings"

// 51_N皇后:
func solveNQueens(n int) [][]string {
	chessboard, ans := make([][]string, n), make([][]string, 0)
	// 构建棋盘
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	// 初始化棋盘
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}
	// 定义回溯函数
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowList := range chessboard {
				temp[i] = strings.Join(rowList, "")
			}
			ans = append(ans, temp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = "Q"
				dfs(row + 1)
				// 恢复现场
				chessboard[row][i] = "."
			}
		}
	}
	dfs(0)
	return ans
}

func isValid(n, row, col int, chessboard [][]string) bool {
	// 判断当前列存在皇后
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 判断斜角45度是否存在皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 判断斜角135度是否存在皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}
