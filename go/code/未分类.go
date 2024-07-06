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
