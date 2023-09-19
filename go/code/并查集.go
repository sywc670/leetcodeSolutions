package code

// 547. 省份数量
// 未掌握
func findCircleNumV3(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
		// error: p = i
		// p是一个拷贝，除非是引用变量，否则不能够影响到原有变量
	}
	var find func(int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(i, j int) {
		parent[find(i)] = find(j)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	for i, p := range parent {
		if p == i {
			ans++
		}
	}
	return
}
