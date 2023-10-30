package code

// lc 239
// 队列存储下标而不是数值，可以保存队首位置信息，
// 从而只使用右指针也可以判断滑动窗口长度是否超过k
func maxSlidingWindow(nums []int, k int) (ans []int) {
	queue := []int{}
	for i, x := range nums {
		// in
		for len(queue) > 0 && x >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// out
		if i-queue[0]+1 > k {
			queue = queue[1:]
		}
		// record
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return
}
