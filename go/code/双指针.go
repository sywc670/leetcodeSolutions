package code

import "strconv"

// 443. 压缩字符串
func compress(chars []byte) (ans int) {
	if len(chars) == 0 {
		return 0
	}
	left, count, index := 0, 0, 0
	// left is actually unnecesary
	summary := func() {
		ans++
		chars[index] = chars[left]
		index++
		if count > 1 {
			countStr := strconv.Itoa(count)
			for _, r := range countStr {
				chars[index] = byte(r)
				index++
			}
			// 处理返回位数
			ans++
			for count/10 >= 1 {
				count /= 10
				ans++
			}
		}
	}
	for index, right := range chars {
		if right == chars[left] {
			count++
		} else {
			summary()
			left, count = index, 1
		}
	}
	// last char remains to be dealt with
	summary()
	return
}

// 392. 判断子序列
func isSubsequence(s string, t string) bool {
	sp, tp := 0, 0
	for sp < len(s) && tp < len(t) {
		if s[sp] == t[tp] {
			sp++
			tp++
			continue
		}
		tp++
	}
	if sp == len(s) && tp <= len(t) {
		return true
	}
	return false
}
