package code

// 思路：单调栈题目，首先有一串数字，当数字是单调减少时，入栈，
// 当遇到比栈顶更大的数字时，出栈之前的一些数字，并做一些工作
// 遍历顺序和单调顺序可能变化

// lc 739
// 单调栈每次都清除无用数据，保持元素有序
// 从右到左
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= t {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}

// 从左到右
func dailyTemperaturesV1(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// lc 42
// 思路：横向计算水面积，需要左边最近的顶和右边最近的顶以及中间的底
// 从左到右遍历，遇到比栈顶大，出栈并计算
func trapV1(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= h { // 当高度连续时也成立，因为后续相同高度面积为0
			bottomHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			width := i - stack[len(stack)-1] - 1
			ans += width * (min(height[stack[len(stack)-1]], h) - bottomHeight)
		}
		stack = append(stack, i)
	}
	return
}
