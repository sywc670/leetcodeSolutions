package code

// lc 560. 和为 K 的子数组
// 暴力枚举
func subarraySum(nums []int, k int) (ans int) {
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return
}

// 未掌握
// 前缀和+哈希表
// 任意连续数组和可以通过两个前缀和的差来表示
// nums[i]+…+nums[j]=prefixSum[j]−prefixSum[i−1]
// 故意让 prefixSum[-1] 为 0，使得通式在i=0时也成立
// 即在遍历之前，map 初始放入 0:1 键值对（前缀和为0出现1次了）
// 边存边查看 map，如果 map 中存在 key 为「当前前缀和 - k」，
// 说明这个之前出现的前缀和，满足「当前前缀和 - 该前缀和 == k」，次数累加。
// 前缀和的值可能会重复，所以需要次数来保存重复个数
func subarraySumV2(nums []int, k int) (ans int) {
	prefixSet := map[int]int{0: 1}
	preSum := 0
	for _, n := range nums {
		preSum += n
		if prefixSet[preSum-k] > 0 {
			ans += prefixSet[preSum-k]
		}
		prefixSet[preSum]++
	}
	return
}
