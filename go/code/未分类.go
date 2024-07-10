package code

// 1679. K 和数对的最大数目
func maxOperations(nums []int, k int) (count int) {
	m := make(map[int]int)
	for _, n := range nums {
		if m[k-n] > 0 {
			count++
			m[k-n]--
		} else {
			m[n]++
		}
	}
	return
}
func maxOperationsOld(nums []int, k int) (count int) {
	m := make(map[int]int)
	for _, n := range nums {
		_, ok := m[n]
		if !ok {
			m[n] = 1
			continue
		}
		m[n]++
	}
	for key := range m {
		for m[key] > 0 {
			if k == 2*key {
				count += m[key] / 2
				break
			}
			peer := k - key
			if _, ok := m[peer]; ok && m[peer] > 0 {
				count++
				m[peer]--
				m[key]--
			} else {
				break
			}
		}
	}
	return
}

// 1456. 定长子串中元音的最大数目
func maxVowels(s string, k int) (ans int) {
	var cnt int
	for i, in := range s {
		if in == 'a' || in == 'e' || in == 'i' || in == 'o' || in == 'u' {
			cnt++
		}
		if i < k-1 {
			continue
		}
		ans = max(ans, cnt)
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			cnt--
		}
	}
	return
}

// 1004. 最大连续1的个数 III
// 下面方法是统计1，统计0的更好
func longestOnes(nums []int, k int) (ans int) {
	left, right := 0, 0
	var ones int
	for ; right < len(nums); right++ {
		l := right - left + 1
		if nums[right] == 1 {
			ones++
		}
		if l <= ones+k {
			ans = max(ans, l)
			continue
		}
		for l > ones+k && l > 1 {
			if nums[left] == 1 {
				ones--
			}
			left++
			l = right - left + 1
		}
	}
	return
}

func longestOnesV1(nums []int, k int) (ans int) {
	left, cnt0 := 0, 0
	for right, x := range nums {
		cnt0 += 1 - x
		for cnt0 > k {
			cnt0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
