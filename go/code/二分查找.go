package code

import (
	"math"
	"slices"
	"sort"
)

// 小技巧：要找第一个满足的值使用>=或<=，找最后一个满足的值可以通过+1-1来变换等式
// 规避边界条件也可以使用辅助函数

func searchRange(nums []int, target int) []int {
	first := bsearch(nums, target)
	if first == len(nums) || nums[first] != target || first < 0 {
		return []int{-1, -1}
	}
	last := bsearch(nums, target+1) - 1
	return []int{first, last}
}
func bsearch(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// 162. 寻找峰值
func findPeakElement(nums []int) (mid int) {
	n := len(nums)
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt
		}
		return nums[i]
	}
	left, right := -1, n
	for {
		mid = (left + right) / 2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return
		}
		if get(mid) < get(mid-1) {
			right = mid
		} else {
			left = mid
		}
	}
}

// lc 153
func findMinV1(nums []int) int {
	n := len(nums)
	lastn := nums[n-1]
	left, right := -1, n
	for left+1 < right {
		mid := left + (right-left)>>1
		// 当mid指向的值大于最后一个数，往右寻找
		if nums[mid] > lastn {
			left = mid
			continue
		}
		// 边界条件判断
		if mid == 0 {
			return nums[mid]
		}
		if mid == n-1 {
			if nums[mid] < nums[mid-1] {
				return nums[mid]
			}
		}
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		} else {
			right = mid
		}
	}
	return -1
}

// 红色表示最小值的左边，蓝色表示最小值及其右边
// 最后一个数必然是蓝色，故不需要判断，判断了也无所谓
func findMinV2(nums []int) int {
	n := len(nums)
	left, right := -1, n-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[n-1] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[right]
}

// lc 33
func search(nums []int, target int) int {
	n := len(nums)
	left, right := -1, n-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[n-1] {
			left = mid
		} else {
			right = mid
		}
	}
	if target <= nums[n-1] {
		l, r := right-1, n
		for l+1 < r {
			m := l + (r-l)/2
			if nums[m] > target {
				r = m
			} else if nums[m] == target {
				return m
			} else {
				l = m
			}
		}
	} else {
		l, r := -1, right
		for l+1 < r {
			m := l + (r-l)/2
			if nums[m] > target {
				r = m
			} else if nums[m] == target {
				return m
			} else {
				l = m
			}
		}
	}
	return -1
}

func guess(int) int {
	return 1
}

// 374. 猜数字大小
func guessNumber(n int) int {
	left, right := -1, n
	for left+1 < right {
		mid := left + (right-left)>>1
		res := guess(mid)
		if res == 1 {
			left = mid
		} else if res == -1 {
			right = mid
		} else {
			return mid
		}
	}
	return n
}

// lc 2300. 咒语和药水的成功对数
func successfulPairs(spells []int, potions []int, success int64) (res []int) {
	slices.Sort(potions)
	res = make([]int, len(spells))
	for i, s := range spells {
		x := int(math.Ceil(float64(success) / float64(s))) // 向上取整
		res[i] = len(potions) - sort.SearchInts(potions, x)
	}
	return
}

// 704. 二分查找
func searchV(nums []int, target int) int {
	left, right := -1, len(nums)

	for left+1 < right {
		mid := left + (right-left)>>1

		if nums[mid] > target {
			right = mid
			continue
		}
		if nums[mid] < target {
			left = mid
			continue
		}

		return mid
	}

	return -1
}

// 875. 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	left := 0                  // 恒为 false
	right := slices.Max(piles) // 恒为 true
	for left+1 < right {       // 开区间不为空
		mid := (left + right) / 2
		sum := n
		for _, p := range piles {
			sum += (p - 1) / mid
		}
		if sum <= h {
			right = mid // 循环不变量：恒为 true
		} else {
			left = mid // 循环不变量：恒为 false
		}
	}
	return right // 最小的 true
}

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
