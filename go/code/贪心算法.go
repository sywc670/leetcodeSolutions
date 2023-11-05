package code

import "sort"

// lc 455
func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return
}

// lc 121 买卖股票的最佳时机
func maxProfitV8(prices []int) (ans int) {
	minP := prices[0]
	for _, p := range prices {
		if p > minP {
			ans = max(ans, p-minP)
			continue
		}
		minP = p
	}
	return
}

// lc 55 跳跃游戏
// 这里维护的是能量值，或者是距离，但是可以使用绝对位置来简化
func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	energy := 0
	for i := 0; i < n-1; i++ {
		energy = max(nums[i], energy)
		if energy == 0 {
			return false
		}
		if energy+i >= n-1 {
			return true
		}
		energy--
	}
	return false
}
