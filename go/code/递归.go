package code

// lc 1071：字符串的最大公因子
func gcdOfStrings(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	if m == n {
		if str1 == str2 {
			return str1
		}
		return ""
	}
	minLength := min(m, n)
	for i := 0; i < minLength; i++ {
		if str1[i] != str2[i] {
			return ""
		}
	}
	if m > n {
		return gcdOfStrings(str1[n:], str2)
	}
	return gcdOfStrings(str1, str2[m:])
}
