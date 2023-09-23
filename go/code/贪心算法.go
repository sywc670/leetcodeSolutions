package code

import "sort"

// lc 455
func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return
}
