package code

import "math"

func unbound_knapsack(c int, w []int, v []int) int {
	n := len(w)
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			return 0
		}
		if c < w[i] {
			return dfs(i-1, c)
		}
		return max(dfs(i-1, c), dfs(i, c-w[i])+v[i])
	}
	return dfs(n-1, c)
}

// leetcode 322
func coinChange(coins []int, amount int) (res int) {
	n := len(coins)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, c int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		C := &cache[i][c]
		if *C != -1 {
			return *C
		}
		defer func() { *C = res }()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	res = dfs(n-1, amount)
	if res < math.MaxInt/2 {
		return res
	}
	return -1
}
