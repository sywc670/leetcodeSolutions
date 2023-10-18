package code

// lc 1
// 与167题区别在于无序，且要返回下标所以不能排序
// 思路在于两个数相加变成nums[i] == target - nums[j]
// 遍历每个数可以存下之前数的值和下标，之后每个数再查找需要的数是否存在
func twoSumV(nums []int, target int) []int {
	index := make(map[int]int)
	for i, x := range nums {
		if j, ok := index[target-x]; ok {
			return []int{j, i}
		}
		index[x] = i
	}
	return nil
}
