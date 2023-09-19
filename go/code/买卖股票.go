package code

import (
	"math"
)

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
				// 不可能发生的情况赋值负无穷
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

// dp
func maxProfitV1(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	dp[0][1] = math.MinInt
	for i, v := range prices {
		dp[i+1][1] = max(dp[i][1], dp[i][0]-v)
		dp[i+1][0] = max(dp[i][0], dp[i][1]+v)
	}
	return dp[n][0]
}

// one array or two var
func maxProfitV5(prices []int) int {
	dp := [2]int{}
	dp[1] = math.MinInt
	for _, v := range prices {
		dp[1], dp[0] = max(dp[1], dp[0]-v), max(dp[0], dp[1]+v)
	}
	return dp[0]
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

// dp
func maxProfitV4(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+2)
	dp[1][1] = math.MinInt
	for i, v := range prices {
		dp[i+2][1] = max(dp[i+1][1], dp[i][0]-v)
		dp[i+2][0] = max(dp[i+1][0], dp[i+1][1]+v)
	}
	return dp[n+1][0]
}

// space optimization
func maxProfitV6(prices []int) int {
	dp := [2]int{}
	pre := 0
	dp[1] = math.MinInt
	for _, v := range prices {
		// 当存在一天冷冻期时，每次循环保存这一次的初始dp[0]给一个全局变量，下一次赋值即可
		// 当存在两天冷冻期时，可以再加一个变量，如下
		// pre0, pre1, dp[1], dp[0] = pre1, dp[0], max(dp[1], pre0-v), max(dp[0], dp[1]+v)
		pre, dp[1], dp[0] = dp[0], max(dp[1], pre-v), max(dp[0], dp[1]+v)
	}
	return dp[0]
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

// dp
func maxProfitV7(k int, prices []int) int {
	dp := make([][2]int, k+2)
	for j := 1; j <= k+1; j++ {
		dp[j][1] = math.MinInt / 2
	}
	dp[0][0] = math.MinInt / 2
	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+p)
			dp[j][1] = max(dp[j][1], dp[j-1][0]-p)
		}
	}
	return dp[k+1][0]
}

// 714. 买卖股票的最佳时机含手续费
func maxProfitV9(prices []int, fee int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2 // 防止溢出
			}
			return
		}
		p := &memo[i][hold]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if hold == 1 {
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		}
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i]-fee)
	}
	return dfs(n-1, 0)
}
