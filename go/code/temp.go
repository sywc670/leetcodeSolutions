package code

func isSubsequence(s string, t string) bool {
	sp, tp := 0, 0
	for sp < len(s) && tp < len(t) {
		if s[sp] == t[sp] {
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
