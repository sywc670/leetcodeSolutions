package code

import "math"

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
func robV2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n+2)
	dp[0], dp[1] = 0, 0
	for i, v := range nums {
		dp[i+2] = max(dp[i]+v, dp[i+1])
	}
	return dp[n+1]
}

// 递推优化
func robV3(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	f0, f1 := 0, 0
	var f int
	for _, v := range nums {
		f = max(f0+v, f1)
		f0 = f1
		f1 = f
	}
	return f
}

// lc 516
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lc 543
// 二叉树求最大值类型的题可能会用到遍历每个节点比较存储最大值的方法
// 在顶点求链的值用加法，而子树求链的值用max比较，因为子树的链要过顶点必须不能分叉
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode) int
	dfs = func(t *TreeNode) int {
		if t == nil {
			return -1
		}
		l := dfs(t.Left)
		r := dfs(t.Right)
		// return value is depth
		ans = max(ans, l+r+2)
		return max(l, r) + 1
	}
	dfs(root)
	return ans
}

// lc 124
// If the result of subtree is below zero
// return zero because that result is not useful
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(t *TreeNode) (res int) {
		if t == nil {
			return 0
		}
		l := dfs(t.Left)
		r := dfs(t.Right)
		res = max(l, r) + t.Val
		ans = max(ans, l+r+t.Val)
		if res < 0 {
			return 0
		}
		return
	}
	dfs(root)
	return ans
}

// lc 2246
func longestPath(parent []int, s string) (ans int) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		pa := parent[i]
		g[pa] = append(g[pa], i)
	}

	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range g[x] {
			len := dfs(y) + 1
			if s[y] != s[x] {
				ans = max(ans, maxLen+len)
				maxLen = max(maxLen, len)
			}
		}
		return
	}
	dfs(0)
	return ans + 1
}

// lc 337
func robV1(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lRob, lNotRob := dfs(node.Left)
		rRob, rNotRob := dfs(node.Right)
		notRob := max(lRob, lNotRob) + max(rRob, rNotRob)
		return lNotRob + rNotRob + node.Val, notRob
	}
	return max(dfs(root))
}

// lc 1143 最长公共子序列
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
func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(i int) int
	dfs = func(i int) (res int) {
		if i > n-1 {
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
