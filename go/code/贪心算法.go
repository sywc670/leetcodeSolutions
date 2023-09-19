package code

import (
	"math"
	"slices"
	"sort"
)

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

// 435. 无重叠区间
// 只能按照左端点从大到小排序，或者右端点从小到大排序，原理：
// 固定右端点从小到大后，所有后续区间只需要比较当前区间左端点和之前区间最大右端点，
// 只要当前区间的左端点超过了之前区间最大右端点，就说明不重叠，反之重叠，不计入当前区间
func eraseOverlapIntervalsV1(intervals [][]int) int {
	// 排序区间
	n := len(intervals)
	if n < 1 {
		return 0
	}
	slices.SortFunc(intervals, func(i, j []int) int { return i[1] - j[1] })

	cnt := 1
	maxRight := intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] >= maxRight {
			cnt++
			maxRight = v[1]
		}
	}
	return n - cnt
}

// 452. 用最少数量的箭引爆气球
// 按右端排序，设立标杆，如果左端小于标杆，不需要更多弓箭，如果大于，需要多一只弓箭
func findMinArrowShots(points [][]int) (cnt int) {
	slices.SortFunc(points, func(i, j []int) int { return i[1] - j[1] })

	overlap := math.MinInt
	for _, p := range points {
		if p[0] > overlap {
			cnt++
			overlap = p[1]
		}
	}
	return
}
