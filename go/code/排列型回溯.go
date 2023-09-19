package code

import "strings"

// lc 46
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// Thought: 新建一个visited结构保存是否访问过，依次枚举
func permute(nums []int) (ans [][]int) {
	n := len(nums)
	res := make([]int, n)
	visited := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= n {
			ans = append(ans, append([]int{}, res...))
			return
		}
		for k, v := range nums {
			if !visited[k] {
				visited[k] = true
				res[i] = v
				dfs(i + 1)
				visited[k] = false
			}
		}
	}
	dfs(0)
	return
}

// lc 51
// 思路：i为横坐标，j为纵坐标，从上往下顺序下棋子，每次递归完成纵坐标的选取，并保存
func solveNQueens(n int) (ans [][]string) {
	col := make([]int, n)
	visited := make([]bool, n)
	// diag用下标保存斜线的固定值
	diag1 := make([]bool, 3*n-1)
	diag2 := make([]bool, 2*n-1)
	abs := func(x int) int {
		if x < 0 {
			return -x + n // 负数范围由正数范围的延长来表示
		}
		return x
	}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			board := make([]string, n)
			for r, c := range col {
				// 由于输出要一行一行输出，故N皇后算法最好也按照从上到下确定坐标，好拿到每一行纵坐标
				board[r] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
			}
			ans = append(ans, board)
			return
		}
		for j := 0; j < n; j++ {
			// 不在同一列同一斜线
			if !visited[j] && !diag1[abs(i-j)] && !diag2[i+j] {
				visited[j], diag1[abs(i-j)], diag2[i+j] = true, true, true
				col[i] = j
				dfs(i + 1)
				visited[j], diag1[abs(i-j)], diag2[i+j] = false, false, false
			}
		}
	}
	dfs(0)
	return
}

// lc 52
func totalNQueens(n int) (ans int) {
	col := make([]int, n)
	visited := make([]bool, n)
	// diag用下标保存斜线的固定值
	diag1 := make([]bool, 3*n-1)
	diag2 := make([]bool, 2*n-1)
	abs := func(x int) int {
		if x < 0 {
			return -x + n // 负数范围由正数范围的延长来表示
		}
		return x
	}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans += 1
			return
		}
		for j := 0; j < n; j++ {
			// 不在同一列同一斜线
			if !visited[j] && !diag1[abs(i-j)] && !diag2[i+j] {
				visited[j], diag1[abs(i-j)], diag2[i+j] = true, true, true
				col[i] = j
				dfs(i + 1)
				visited[j], diag1[abs(i-j)], diag2[i+j] = false, false, false
			}
		}
	}
	dfs(0)
	return
}
