package code

func zero_one_knapsack(c int, w []int, v []int) int {
	n := len(w)
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			return 0
		}
		if c < w[i] {
			return dfs(i-1, c)
		}
		return max(dfs(i-1, c), dfs(i-1, c-w[i])+v[i])
	}
	return dfs(n-1, c)
}

// leetcode 494 目标和
func findTargetSumWays(nums []int, target int) int {
	// pos == (sum + target) / 2
	n := len(nums)
	var s int
	for _, n := range nums {
		s += n
	}
	p := s + target
	if p%2 == 1 || p < 0 {
		return 0
	}
	p = p / 2

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, p+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, p int) int
	dfs = func(i, p int) (res int) {
		if i < 0 {
			if p != 0 {
				return 0
			} else {
				return 1
			}
		}
		c := &cache[i][p]
		if *c != -1 {
			return *c
		}
		defer func() {
			*c = res
		}()
		if p < nums[i] {
			return dfs(i-1, p)
		}
		return dfs(i-1, p) + dfs(i-1, p-nums[i])
	}
	return dfs(n-1, p)
}

// 动态规划写法
func findTargetSumWaysV2(nums []int, target int) int {
	// pos == (sum + target) / 2
	n := len(nums)
	var s int
	for _, n := range nums {
		s += n
	}
	p := s + target
	if p%2 == 1 || p < 0 {
		return 0
	}
	p = p / 2

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, p+1)
	}
	dp[0][0] = 1
	for i := range nums {
		for c := 0; c <= p; c++ {
			if c < nums[i] {
				dp[i+1][c] = dp[i][c]
				continue
			}
			dp[i+1][c] = dp[i][c] + dp[i][c-nums[i]]
		}
	}
	return dp[n][p]
}

// 空间优化
func findTargetSumWaysV3(nums []int, target int) int {
	// pos == (sum + target) / 2
	n := len(nums)
	var s int
	for _, n := range nums {
		s += n
	}
	p := s + target
	if p%2 == 1 || p < 0 {
		return 0
	}
	p = p / 2

	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, p+1)
	}
	dp[0][0] = 1
	for i := range nums {
		for c := 0; c <= p; c++ {
			if c < nums[i] {
				dp[(i+1)%2][c] = dp[i%2][c]
				continue
			}
			dp[(i+1)%2][c] = dp[i%2][c] + dp[i%2][c-nums[i]]
		}
	}
	return dp[n%2][p]
}

func findTargetSumWaysV4(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	f := make([]int, target+1)
	f[0] = 1
	for _, x := range nums {
		for c := target; c >= x; c-- {
			// note: 这里是倒序，f(i+1,c)依赖于f(i,c-x)，而正序时的f(i,c-x)已经被f(i+1,c-x)给替换了，无法使用
			// 所以计算出来结果不正确，倒序则不会出现问题
			f[c] += f[c-x]
		}
	}
	return f[target]
}
