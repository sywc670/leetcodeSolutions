package code

// 小技巧：滑动窗口可以用range来移动右指针

// 209. 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	left, sum, count := 0, 0, 0
	for _, v := range nums {
		sum += v
		if sum < target {
			continue
		}

		for sum-nums[left] > target {
			sum -= nums[left]
			left++
		}
	}

	return count
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
		for set[byte(current)] {
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

// 438. 找到字符串中所有字母异位词
// 提示：数组可以直接比较
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
