package code

import "slices"

// lc 198
// 记忆化搜索/回溯
func rob(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i < 0 {
			return 0
		}
		C := &cache[i]
		if *C != -1 {
			return *C
		}
		defer func() {
			*C = res
		}()
		return max(dfs(i-2)+nums[i], dfs(i-1))
	}
	return dfs(n - 1)
}

// 递推
func robV1(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i, v := range nums {
		dp[i+2] = max(dp[i]+v, dp[i+1])
	}
	return dp[n+1]
}

// 递推优化
func robV2(nums []int) int {
	f0, f1 := 0, 0
	for _, v := range nums {
		f := max(f0+v, f1)
		f0, f1 = f1, f
	}
	return f1
}

// 1143 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return 0
		}
		c := &cache[i][j]
		if *c != -1 {
			return *c
		}
		defer func() {
			*c = res
		}()
		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(m-1, n-1)
}

// dp
func longestCommonSubsequenceV1(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, t1 := range text1 {
		for j, t2 := range text2 {
			if t1 == t2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}

// optimize
func longestCommonSubsequenceV2(text1, text2 string) int {
	n := len(text2)
	dp := make([]int, n+1)
	for _, t1 := range text1 {
		pre := 0 // 初始值其实就是dp[0]，本题中pre的值与i无关，所以固定初始为0
		// 这里使用pre是因为dp[i-1][j-1]会被覆盖掉，但也不用倒序，因为j-1距离j的值是固定的且为1
		// 零一背包中是c-w[i]，距离是不固定的，所以最好使用倒序
		for j, t2 := range text2 {
			if t1 == t2 {
				// 下面的赋值是语法糖，前面的赋值不会影响后面的赋值，因为有临时变量
				// 如下：
				// tmp := dp[j+1]
				// dp[j+1] = pre+1
				// pre = tmp
				dp[j+1], pre = pre+1, dp[j+1]
			} else {
				pre = dp[j+1]
				dp[j+1] = max(dp[j], dp[j+1])
			}
		}
	}
	return dp[n]
}

// lc 72
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		c := &cache[i][j]
		if *c != -1 {
			return *c
		}
		defer func() {
			*c = res
		}()
		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		}
		return min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
	}
	return dfs(m-1, n-1)
}

// dp
func minDistanceV1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 转化为dp时，i、j的取值范围需要注意，由于为了不取零值，i、j范围分别为[1,m],[1,n]
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i, w1 := range word1 {
		dp[i+1][0] = i + 1
		for j, w2 := range word2 {
			if w1 == w2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i+1][j], dp[i][j+1], dp[i][j]) + 1
			}
		}
	}
	return dp[m][n]
}

// one slice dp
func minDistanceV2(word1 string, word2 string) int {
	n := len(word2)
	dp := make([]int, n+1)
	// 转化为dp时，i、j的取值范围需要注意，由于为了不取零值，i、j范围分别为[1,m],[1,n]
	for j := 1; j <= n; j++ {
		dp[j] = j
	}
	for _, w1 := range word1 {
		// pre的值代表了word2为空后word1剩余i个字符的情况，从零将i增加
		pre := dp[0]
		dp[0]++
		for j, w2 := range word2 {
			if w1 == w2 {
				dp[j+1], pre = pre, dp[j+1]
			} else {
				dp[j+1], pre = min(dp[j], dp[j+1], pre)+1, dp[j+1]
			}
		}
	}
	return dp[n]
}

// lc 300
// 思路：1.枚举 子问题是以nums[i]结尾的子序列最长长度，枚举之前的所有nums[j]结尾的子序列，
// 只要一个入参，选择枚举来做
// 2.选与不选 会比较两个值的大小，所以需要两个入参
// 3.这题也可以用单调栈+二分查找做，见对应代码
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(i int) int
	// 返回值保存长度，不能使用全局变量，有嵌套函数会对变量进行修改
	dfs = func(i int) (res int) {
		if i > n-1 { // 倒着枚举可以不写这句
			return 0
		}
		if memo[i] > 0 {
			return memo[i]
		}
		defer func() {
			memo[i] = res
		}()
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				res = max(res, dfs(j))
			}
		}
		res++
		return res
	}
	for i := range nums {
		ans = max(ans, dfs(i))
	}
	return
}

// dp
// note: 未完全掌握
func lengthOfLISV1(nums []int) (ans int) {
	n := len(nums)
	dp := make([]int, n)
	for i := range nums {
		for j := range nums[:i] {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
		ans = max(ans, dp[i])
	}
	return
	// return dp[n-1]
	// SOLVE: dp[n-1]不一定最大，因为dp[i] = max(dp[j])+1，而j<i && nums[j]<nums[i]
	// 等式右边没有取所有比i小的数，所以不一定最大
	// 另外dp[n-1]意为取最后一位时的最大长度，不一定最大
}

// lc 70. 爬楼梯
func climbStairs(n int) int {
	prev := 1
	cur := 1
	for i := 2; i < n+1; i++ {
		temp := cur
		cur = prev + cur
		prev = temp
	}
	return cur
}

// 1137. 第 N 个泰波那契数
func tribonacci(n int) int {
	t0, t1, t2 := 0, 1, 1
	if n == 0 {
		return t0
	}
	if n == 1 || n == 2 {
		return 1
	}
	for i := 0; i < n-2; i++ {
		newT := t0 + t1 + t2
		t0, t1, t2 = t1, t2, newT
	}
	return t2
}

// 746. 使用最小花费爬楼梯
// 原问题：n阶台阶上完之后最小cost，用dfs(n)表示
// 子问题：在第i阶时最小cost
// 当前操作：如果从i-1来，那么dfs(i) = dfs(i-1) + cost[i-1]，如果i-2来，那么dfs(i) = dfs(i-2) + cost[i-2]
func minCostClimbingStairs(cost []int) (ans int) {
	var dfs func(i int) int
	cache := make([]int, len(cost)+1)
	for i := range cache {
		cache[i] = -1
	}
	dfs = func(i int) int {
		if i == 0 || i == 1 {
			return 0
		}
		v := cache[i]
		if v != -1 {
			return v
		}
		res := min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		cache[i] = res
		return res
	}
	return dfs(len(cost))
}

// 435. 无重叠区间
// dp会超时
func eraseOverlapIntervals(intervals [][]int) int {
	// 排序区间
	n := len(intervals)
	if n < 1 {
		return 0
	}
	slices.SortFunc(intervals, func(i, j []int) int { return i[0] - j[0] })

	// dp[i] = max(dp[j])+1 j<i && rj < li
	// 以第i个为最后一个区间，dp[i]表示最大的不重叠区间数
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[j]) + 1
			}
		}
	}

	maxdp := dp[0]
	for _, v := range dp {
		if v > maxdp {
			maxdp = v
		}
	}

	return n - maxdp
}

// 62. 不同路径
// dp[i+1][j+1] = d[i][j+1]+d[i+1][j]
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[m-1][n-1]
}

// 790. 多米诺和托米诺平铺
func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1], f[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = (f[i-1]*2 + f[i-3]) % (1e9 + 7)
	}
	return f[n]
}
