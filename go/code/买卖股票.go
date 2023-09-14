package code

import "math"

// lc 122
func maxProfit(prices []int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return
		}
		p := &memo[i][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0)
}

// lc 309
func maxProfitV2(prices []int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt
			}
			return
		}
		p := &memo[i][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-2, 0)-prices[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0)
}

// lc 188
func maxProfitV3(k int, prices []int) int {
	n := len(prices)
	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, k+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1} // -1 表示还没有计算过
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, k, hold int) (res int) {
		if k < 0 {
			return math.MinInt / 2
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2
			}
			return
		}
		p := &memo[i][k][hold]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if hold == 1 {
			return max(dfs(i-1, k, 1), dfs(i-1, k-1, 0)-prices[i])
		}
		return max(dfs(i-1, k, 0), dfs(i-1, k, 1)+prices[i])
	}
	return dfs(n-1, k, 0)
}
