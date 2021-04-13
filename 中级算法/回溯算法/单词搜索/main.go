package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCESEEEFS"))
}

func exist(board [][]byte, word string) bool {

	// 初始化,记录走过的位置
	visit := map[int]bool{}

	var dfs func(r, c, index int) bool

	dfs = func(r, c, index int) bool {

		if index == len(word) {
			return true
		}
		// 越界
		if !overstep(r, c, len(board), len(board[0])) {
			return false
		}

		// 不等于word
		if board[r][c] != word[index] {
			return false
		}

		// 访问过的单元格
		if visit[r*10+c] {
			return false
		}

		visit[r*10+c] = true
		defer func() {
			// 回溯复原visit(清除脚步)
			visit[r*10+c] = false
		}()
		return dfs(r-1, c, index+1) || dfs(r+1, c, index+1) || dfs(r, c-1, index+1) || dfs(r, c+1, index+1)
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			// 减少遍历次数
			if board[r][c] == word[0] {
				if dfs(r, c, 0) {
					return true
				}
			}
		}
	}

	return false
}

func overstep(r, c, n, m int) bool {
	if 0 <= r && r < n && 0 <= c && c < m {
		return true
	}
	return false
}
