package code

// 17. 电话号码的字母组合
var mapping = [...]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}
	path := make([]byte, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i > n-1 {
			ans = append(ans, string(path))
			return
		}
		for _, r := range mapping[digits[i]-'0'] {
			path[i] = byte(r)
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

// lc 78
// 选与不选思路
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	if n == 0 {
		return
	}
	path := make([]int, 0, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			// 下面这种写法不行，因为这里添加的path是[]int，是引用变量，
			// 所以会存储指针，而不是值，之后修改path会影响ans
			// ans = append(ans, path)
			return
		}
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
		dfs(i + 1)
	}
	dfs(0)
	return
}

// 枚举思路
func subsetsV2(nums []int) (ans [][]int) {
	n := len(nums)
	path := make([]int, 0, n)
	var dfs func(i int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...))
		if i == n {
			return
		}
		for j := i; j < n; j++ {
			// j是选中的符号下标
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}

// lc 131
// 思路：枚举第x个位置是否会隔开，枚举的字符串开头为i，结尾为j，判断是否回文
// Option: 如果用选与不选来做需要给dfs额外传入开始位置下标
func partition(s string) (ans [][]string) {
	n := len(s)
	cur := make([]string, 0, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, cur...))
			return
		}
		for j := i; j < n; j++ {
			// j是最后位置的指针
			var isPalindrome = func(i, j int) bool {
				for i < j {
					if s[i] != s[j] {
						return false
					}
					i++
					j--
				}
				return true
			}
			if isPalindrome(i, j) == true {
				cur = append(cur, s[i:j+1])
				dfs(j + 1)
				cur = cur[:len(cur)-1]
			}
		}

	}
	dfs(0)
	return
}
