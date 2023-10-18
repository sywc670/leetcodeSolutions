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

// lc 283
func moveZeroes(nums []int) {
	count := 0
	for i, n := range nums {
		if n != 0 {
			nums[count] = n
			if count != i {
				nums[count] = 0
			}
			count++
		}
	}
}

// lc 438
// 提示：数组和切片可以直接比较
func findAnagrams(s string, p string) (ans []int) {
	m, n := len(s), len(p)
	if n > m {
		return nil
	}
	left, right := 0, -1
	sCount, pCount := [26]int{}, [26]int{}
	for i, r := range p {
		pCount[r-'a']++
		sCount[s[i]-'a']++
		right++
	}
	for right < m {
		if sCount == pCount {
			ans = append(ans, left)
		}
		sCount[s[left]-'a']--
		left++
		right++
		if right < m {
			sCount[s[right]-'a']++
		}
	}
	return
}
