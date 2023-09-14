package code

// 小技巧：滑动窗口可以用range来移动右指针

// lc 209
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, s, left := n+1, 0, 0
	for right, x := range nums {
		s += x
		for s >= target { // 满足要求
			ans = min(ans, right-left+1)
			s -= nums[left]
			left++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}

// lc 713
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k <= 1 {
		return
	}
	p, left := 1, 0
	for right := range nums {
		p *= nums[right]
		for p >= k && left < right {
			p /= nums[left]
			left++
		}
		if p < k {
			ans += right - left + 1
		}
	}
	return
}

// lc 3
func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	set := make(map[byte]bool)
	for right, current := range s {
		for set[byte(current)] == true {
			set[byte(s[left])] = false
			left++
		}
		set[byte(current)] = true
		ans = max(ans, right-left+1)
	}
	return
}
