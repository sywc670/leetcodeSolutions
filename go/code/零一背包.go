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

// leetcode 494
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
