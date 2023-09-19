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
func coinChange(coins []int, amount int) (ans int) {
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
		defer func() {
			*C = res
		}()
		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans = dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func coinChangeV1(coins []int, amount int) (ans int) {
	n := len(coins)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[0][j] = math.MaxInt / 2
		}
	}
	dp[0][0] = 0
	for i := range coins {
		for c := 0; c <= amount; c++ {
			if c < coins[i] {
				dp[(i+1)%2][c] = dp[i%2][c]
			} else {
				dp[(i+1)%2][c] = min(dp[i%2][c], dp[(i+1)%2][c-coins[i]]+1)
			}
		}
	}
	ans = dp[n%2][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func coinChangeV2(coins []int, amount int) (ans int) {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt / 2
	}
	dp[0] = 0
	for _, x := range coins {
		for c := x; c <= amount; c++ {
			// 这里不像零一背包一样使用倒序，因为不会替换之后要使用的值
			dp[c] = min(dp[c], dp[c-x]+1)
		}
	}
	ans = dp[amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
