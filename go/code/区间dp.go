package code

// lc 516
// 思路：可以求s和反转s的最长公共子序列
// 区间dp 选择使用选与不选来做
func longestPalindromeSubseq(s string) int {
	n := len(s)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		C := &cache[i][j]
		if *C != -1 {
			return *C
		}
		defer func() {
			*C = res
		}()
		if s[i] == s[j] {
			return dfs(i+1, j-1) + 2
		}
		return max(dfs(i+1, j), dfs(i, j-1))
	}
	return dfs(0, n-1)
}

// dp
// note: 未完全掌握
func longestPalindromeSubseqV1(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		// i倒序遍历，因为需要先得到i+1
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			// j从i+1正向遍历，使得j-1规避了负数，同时还使得i+1规避了上限
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// one slice
func longestPalindromeSubseqV2(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		// i倒序遍历，因为需要先得到i+1
		pre := 0
		// pre初始值为dp[i+1][j-1]，因为i是倒序，j是正序，并且由于j初始值为i+1，故简化dp[i+1][i]
		// dp[i+1][i]已经不满足i<j的条件，必定找不到回文子序列
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			// j从i+1正向遍历，使得j-1规避了负数，同时还使得i+1规避了上限
			if s[i] == s[j] {
				pre, dp[j] = dp[j], pre+2
			} else {
				pre = dp[j]
				dp[j] = max(dp[j], dp[j-1])
			}
		}
	}
	return dp[n-1]
}
