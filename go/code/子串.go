package code

// lc 14
func longestCommonPrefix(strs []string) (ans string) {
	n := 201
	count := len(strs)
	for i := 0; i < count; i++ {
		n = min(n, len(strs[i]))
	}
	for i := 0; i < n; i++ {
		for _, str := range strs {
			if str[i] != strs[0][i] {
				return
			}
		}
		ans += string(strs[0][i])
	}
	return
}
