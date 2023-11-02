package code

import (
	"slices"
	"sort"
)

// lc 581
func findUnsortedSubarray(nums []int) int {
	if slices.IsSorted(nums) {
		return 0
	}
	newNums := append([]int{}, nums...)
	sort.Ints(newNums)
	left, right := 0, len(nums)-1
	for ; newNums[left] == nums[left]; left++ {
	}
	for ; newNums[right] == nums[right]; right-- {
	}
	return right - left + 1
}
