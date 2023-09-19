package code

// lc 724. 寻找数组的中心下标
// 思路是前缀和
// 其它思路：1. 维持前缀和与后缀和，不用考虑端点
// 2. leetcode官方解法
func pivotIndex(nums []int) int {
	prefixSum := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		prefixSum[i] = sum
	}
	// error: 如果数组的和为0，中心下标为0
	// solve: 中心下标不计入左右数组和中
	if sum-prefixSum[0] == 0 {
		return 0
	}
	// 已经排除两个端点
	for i := 1; i < len(nums)-1; i++ {
		if prefixSum[i-1]+prefixSum[i] == sum { // 核心代码
			return i
		}
	}
	// 判断最后一个数是否为中心下标
	if sum-nums[len(nums)-1] == 0 {
		return len(nums) - 1
	}
	return -1
}

// 1732. 找到最高海拔
func largestAltitude(gain []int) int {
	sum := 0
	top := 0
	for _, delta := range gain {
		sum += delta
		top = max(top, sum)
	}
	return top
}

// 238. 除自身以外数组的乘积
// solve: 定义好L\R具体代表什么，L[i]代表i左侧数组乘积和，不包含i
func productExceptSelf(nums []int) []int {
	n := len(nums)
	L, R := make([]int, n), make([]int, n)
	L[0], R[n-1] = 1, 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = L[i] * R[i]
	}
	return ans
}

// 303. 区域和检索 - 数组不可变
type NumArray struct {
	nums []int
	sums map[int]int
}

func Constructor_1(nums []int) NumArray {
	na := NumArray{
		nums: nums,
		sums: map[int]int{-1: 0},
	}

	for i, j := range nums {
		na.sums[i] = na.sums[i-1] + j
	}
	return na
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.sums[right] - na.sums[left-1]
}

// TODO:
// lc 560. 和为 K 的子数组
// 题意：有几种 i、j 的组合，使得从第 i 到 j 项的子数组和等于 k。
// 转化为有几种 i、j 的组合，满足 prefixSum[j] - prefixSum[i - 1] == k
// 前缀和+哈希表
// 任意连续数组和可以通过两个前缀和的差来表示，前缀和专门解决这种连续数组求和的问题
// nums[i]+…+nums[j]=prefixSum[j]−prefixSum[i−1]
// 故意让 prefixSum[-1] 为 0，使得通式在i=0时也成立
// 即在遍历之前，map 初始放入 0:1 键值对（前缀和为0出现1次了）
// 边存边查看 map，如果 map 中存在 key 为「当前前缀和 - k」，
// 说明这个之前出现的前缀和，满足「当前前缀和 - 该前缀和 == k」，次数累加。
// 前缀和的值可能会重复，所以需要次数来保存重复个数
func subarraySum(nums []int, k int) (ans int) {
	// when len(nums) == 1, prefixSumCnt[0] must have prefixSumCnt[-1] to sub
	prefixSumCnt := map[int]int{0: 1}
	prefixSum := 0
	for _, n := range nums {
		prefixSum += n
		ans += prefixSumCnt[prefixSum-k]
		prefixSumCnt[prefixSum]++
	}
	return
}

// 437. 路径总和 III
func pathSum(root *TreeNode, targetSum int) (ans int) {
	prefixSumCnt := map[int]int{0: 1}
	var dfs func(node *TreeNode, prefixSum int)
	dfs = func(node *TreeNode, prefixSum int) {
		if node == nil {
			return
		}
		prefixSum += node.Val
		ans += prefixSumCnt[prefixSum-targetSum]
		prefixSumCnt[prefixSum]++

		dfs(node.Left, prefixSum)
		dfs(node.Right, prefixSum)

		prefixSumCnt[prefixSum]--
	}
	dfs(root, 0)
	return
}
