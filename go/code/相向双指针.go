package code

import "sort"

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

// lc 42
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
