package code

import (
	"math"
	"sort"
)

// lc 167
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	p, q := 0, n-1
	for p < q {
		sum := numbers[p] + numbers[q]
		if sum == target {
			return []int{p + 1, q + 1}
		} else if sum > target {
			q--
		} else {
			p++
		}
	}
	return []int{0, 0}
}

// lc 15
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		j, k := i+1, n-1
		target := -nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}

// lc 16
func threeSumClosest(nums []int, target int) (ans int) {
	delta := math.MaxInt
	sort.Ints(nums)
	n := len(nums)
	if nums[0]+nums[1]+nums[2] >= target {
		return nums[0] + nums[1] + nums[2]
	}
	if nums[n-1]+nums[n-2]+nums[n-3] <= target {
		return nums[n-1] + nums[n-2] + nums[n-3]
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i, x := range nums[:n-2] {
		j, k := i+1, n-1
		for j < k {
			sum := x + nums[j] + nums[k]
			sub := sum - target
			if abs(sub) < abs(delta) {
				delta = sub
			}
			if sub < 0 {
				j++
				continue
			}
			k--
		}
	}
	return delta + target
}

// lc 18
func fourSum(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	if n < 4 {
		return
	}
	for i, x := range nums[:n-3] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if x+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			y := nums[j]
			// for j, y := range nums[i+1 : n-2] {
			// 这里不能用range是因为j的下标会从0开始，而不是从i+1开始
			if j > i+1 && y == nums[j-1] {
				continue
			}
			if x+y+nums[j+1]+nums[j+2] > target {
				break
			}
			if x+y+nums[n-1]+nums[n-2] < target {
				continue
			}
			k, m := j+1, n-1
			for k < m {
				sum := x + y + nums[k] + nums[m]
				if sum > target {
					m--
				} else if sum < target {
					k++
				} else {
					ans = append(ans, []int{x, y, nums[k], nums[m]})
					m--
					for k < m && nums[m] == nums[m+1] {
						m--
					}
					k++
					for k < m && nums[k-1] == nums[k] {
						k++
					}
				}
			}
		}
	}
	return
}

// lc 11
func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		var minHeight = min(height[left], height[right])
		area := (right - left) * minHeight
		if area > ans {
			ans = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

// 42. 接雨水
func trap(height []int) (ans int) {
	l, r := 0, len(height)-1
	L, R := 0, 0
	for l <= r {
		L = max(L, height[l])
		R = max(R, height[r])
		if L < R {
			ans += L - height[l]
			l++
		} else {
			ans += R - height[r]
			r--
		}
	}
	return
}
