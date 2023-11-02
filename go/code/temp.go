package code

import (
	"sort"
)

// lc 面试题 17.08： 马戏团人塔
// 排序，使得题目变成LIS
func bestSeqAtIndex(height []int, weight []int) (ans int) {
	l := len(height)
	type person struct {
		height, weight int
	}
	// 也可以用二维数组来做
	personSlice := make([]person, l)
	for i, h := range height {
		personSlice[i] = person{height: h, weight: weight[i]}
	}
	// 身高升序排序，体重降序排序，其中体重降序是为了解决身高重复情况
	// 体重升序会让身高重复的人被错误计算在内
	sort.Slice(personSlice, func(i, j int) bool {
		if personSlice[i].height == personSlice[j].height {
			return personSlice[i].weight > personSlice[j].weight
		}
		return personSlice[i].height < personSlice[j].height
	})

	// LIS
	// 单调栈+二分查找
	mStack := make([]int, 0, l/2)
	mStack = append(mStack, personSlice[0].weight)
	ans = 1
	for i := 1; i < l; i++ {
		target := personSlice[i].weight
		if target > mStack[len(mStack)-1] {
			mStack = append(mStack, target)
			ans = max(ans, len(mStack))
		} else if target < mStack[len(mStack)-1] {
			left, right := -1, len(mStack)
			for left+1 < right {
				mid := (left + right) / 2
				if target > mStack[mid] { // SOLVE: 二分比较成mid下标了
					left = mid
				} else {
					right = mid
				}
			}
			mStack[right] = target // 替代比target大的第一个数
		}
	}
	return
}
