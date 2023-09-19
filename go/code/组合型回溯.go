package code

// lc 77
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。
// Option: 选与不选做法
func combine(n int, k int) (ans [][]int) {
	res := make([]int, 0, k)
	var dfs func(i int)
	dfs = func(i int) {
		d := k - len(res)
		if d == 0 {
			ans = append(ans, append([]int{}, res...))
			return
		}
		for j := i; j >= d; j-- {
			res = append(res, j)
			dfs(j - 1)
			res = res[:len(res)-1]
		}
	}
	dfs(n)
	return
}

// 216. 组合总和 III
// 这里是选与不选
func combinationSum3(k int, n int) (ans [][]int) {
	res := make([]int, 0, k)
	var dfs func(num, sum int)
	dfs = func(num, sum int) {
		if len(res) == k && sum == 0 {
			ans = append(ans, append([]int{}, res...))
			return
		}
		// 顺序很重要，最后一遍dfs时，num+1可能会大于9，此时应当添加答案，而不应该直接返回
		if num > 9 || sum < 0 {
			return
		}
		if sum-num >= 0 {
			res = append(res, num)
			dfs(num+1, sum-num)
			res = res[:len(res)-1]
		}
		dfs(num+1, sum)
	}
	dfs(1, n)
	return
}

// 枚举做法
func combinationSum3V(k int, n int) (ans [][]int) {
	path := make([]int, 0)
	sum := 0
	var dfs func(int, int)
	dfs = func(pos int, pre int) {
		if len(path) == k && sum == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > n || len(path) > k {
			return
		}

		for num := pre + 1; num <= 9; num++ { // pre+1防止重复数字，递增去选择
			path = append(path, num)
			sum += num
			dfs(pos+1, num)
			sum -= num
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return
}

// lc 22
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 思路：选左括号还是选右
func generateParenthesis(n int) (ans []string) {
	res := make([]byte, 2*n)
	var dfs func(i, r int)
	dfs = func(i, r int) {
		if i >= 2*n {
			ans = append(ans, string(res))
			return
		}
		if r > 0 {
			res[i] = ')'
			dfs(i+1, r-1)
		}
		if 2*n-i-r > 0 {
			res[i] = '('
			dfs(i+1, r+1)
		}
	}
	dfs(0, 0)
	return
}
