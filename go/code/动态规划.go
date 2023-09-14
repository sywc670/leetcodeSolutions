package code

import "math"

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
func rob(root *TreeNode) (ans int) {
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

// lc 1143
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
