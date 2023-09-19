package code

// lc 200. 岛屿数量
// 遍历陆地之后将其改成海洋，不需要多余空间来保存
func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

// 547. 省份数量
// 未掌握
// bfs
func findCircleNum(isConnected [][]int) (ans int) {
	visited := make([]bool, len(isConnected))
	for i, v := range visited {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				visited[from] = true
				for to, c := range isConnected[from] {
					if c == 1 && !visited[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return
}

// dfs
func findCircleNumV2(isConnected [][]int) (ans int) {
	visited := make([]bool, len(isConnected))
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for j, conn := range isConnected[i] {
			if !visited[j] && conn == 1 {
				dfs(j)
			}
		}
	}
	for i, v := range visited {
		if !v {
			ans++
			dfs(i)
		}
	}
	return
}

// 1466. 重新规划路线
// 未掌握
// dfs 用一颗多叉树存储，正向为正数，反向为负数
func minReorder(n int, connections [][]int) (ans int) {
	cset := make(map[int][]int)
	for _, c := range connections {
		cset[c[0]] = append(cset[c[0]], c[1])
		cset[c[1]] = append(cset[c[1]], -c[0])
	}
	var dfs func(i int, p int)
	dfs = func(i int, p int) {
		for _, v := range cset[i] {
			if v != p && -v != p {
				if v > 0 {
					ans++
				} else {
					v = -v
				}
				dfs(v, i)
			}
		}
	}
	// 第二个参数是第一个参数的父节点，用于防止重复访问进入死循环
	// 传入n是因为0没有父节点，所以不需要判断
	dfs(0, n)
	return
}
